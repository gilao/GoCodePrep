### 一、C/C++报错?Golang通过?



我们先看一段代码



```go
package main

func foo(arg_val int)(*int) {

    var foo_val int = 11;
    return &foo_val;
}

func main() {

    main_val := foo(666)

    println(*main_val)
}
```



编译运行



```bash
$ go run pro_1.go 
11
```



竟然没有报错!



了解C/C的小伙伴应该知道,这种情况是一定不允许的,因为 外部函数使用了子函数的局部变量, 理论来说,子函数的`foo_val` 的声明周期早就销毁了才对,如下面的C/C代码



```c
#include <stdio.h>

int *foo(int arg_val) {

    int foo_val = 11;

    return &foo_val;
}

int main()
{
    int *main_val = foo(666);

    printf("%d\n", *main_val);
}
```



编译



```bash
$ gcc pro_1.c 
pro_1.c: In function ‘foo’:
pro_1.c:7:12: warning: function returns address of local variable [-Wreturn-local-addr]
     return &foo_val;
            ^~~~~~~~
```



出了一个警告,不管他,再运行



```plain
$ ./a.out 
段错误 (核心已转储)
```



程序崩溃.



如上C/C++编译器明确给出了警告，foo把一个局部变量的地址返回了；反而高大上的go没有给出任何警告，难道是go编译器识别不出这个问题吗？



### 二、Golang编译器得逃逸分析



​		go语言编译器会自动决定把一个变量放在栈还是放在堆，编译器会做**逃逸分析(escape analysis)**，**当发现变量的作用域没有跑出函数范围，就可以在栈上，反之则必须分配在堆**。
go语言声称这样可以释放程序员关于内存的使用限制，更多的让程序员关注于程序功能逻辑本身。



我们再看如下代码:



```go
package main

func foo(arg_val int) (*int) {

    var foo_val1 int = 11;
    var foo_val2 int = 12;
    var foo_val3 int = 13;
    var foo_val4 int = 14;
    var foo_val5 int = 15;


    //此处循环是防止go编译器将foo优化成inline(内联函数)
    //如果是内联函数，main调用foo将是原地展开，所以foo_val1-5相当于main作用域的变量
    //即使foo_val3发生逃逸，地址与其他也是连续的
    for i := 0; i < 5; i++ {
        println(&arg_val, &foo_val1, &foo_val2, &foo_val3, &foo_val4, &foo_val5)
    }

    //返回foo_val3给main函数
    return &foo_val3;
}


func main() {
    main_val := foo(666)

    println(*main_val, main_val)
}
```



我们运行一下



```bash
$ go run pro_2.go 
0xc000030758 0xc000030738 0xc000030730 0xc000082000 0xc000030728 0xc000030720
0xc000030758 0xc000030738 0xc000030730 0xc000082000 0xc000030728 0xc000030720
0xc000030758 0xc000030738 0xc000030730 0xc000082000 0xc000030728 0xc000030720
0xc000030758 0xc000030738 0xc000030730 0xc000082000 0xc000030728 0xc000030720
0xc000030758 0xc000030738 0xc000030730 0xc000082000 0xc000030728 0xc000030720
13 0xc000082000
```



我们能看到`foo_val3`是返回给main的局部变量, 其中他的地址应该是`0xc000082000`,很明显与其他的foo_val1、2、3、4不是连续的.



我们用`go tool compile`测试一下



```bash
$ go tool compile -m pro_2.go
pro_2.go:24:6: can inline main
pro_2.go:7:9: moved to heap: foo_val3
```



果然,在编译的时候, `foo_val3`具有被编译器判定为逃逸变量, 将`foo_val3`放在堆中开辟.



我们在用汇编证实一下:



```bash
$ go tool compile -S pro_2.go > pro_2.S
```



打开pro_2.S文件, 搜索`runtime.newobject`关键字



```go
 ...
 16     0x0021 00033 (pro_2.go:5)   PCDATA  $0, $0
 17     0x0021 00033 (pro_2.go:5)   PCDATA  $1, $0
 18     0x0021 00033 (pro_2.go:5)   MOVQ    $11, "".foo_val1+48(SP)
 19     0x002a 00042 (pro_2.go:6)   MOVQ    $12, "".foo_val2+40(SP)
 20     0x0033 00051 (pro_2.go:7)   PCDATA  $0, $1
 21     0x0033 00051 (pro_2.go:7)   LEAQ    type.int(SB), AX
 22     0x003a 00058 (pro_2.go:7)   PCDATA  $0, $0
 23     0x003a 00058 (pro_2.go:7)   MOVQ    AX, (SP)
 24     0x003e 00062 (pro_2.go:7)   CALL    runtime.newobject(SB)  //foo_val3是被new出来的
 25     0x0043 00067 (pro_2.go:7)   PCDATA  $0, $1
 26     0x0043 00067 (pro_2.go:7)   MOVQ    8(SP), AX
 27     0x0048 00072 (pro_2.go:7)   PCDATA  $1, $1
 28     0x0048 00072 (pro_2.go:7)   MOVQ    AX, "".&foo_val3+56(SP)
 29     0x004d 00077 (pro_2.go:7)   MOVQ    $13, (AX)
 30     0x0054 00084 (pro_2.go:8)   MOVQ    $14, "".foo_val4+32(SP)
 31     0x005d 00093 (pro_2.go:9)   MOVQ    $15, "".foo_val5+24(SP)
 32     0x0066 00102 (pro_2.go:9)   XORL    CX, CX
 33     0x0068 00104 (pro_2.go:15)  JMP 252
 ...
```



看出来, foo_val3是被runtime.newobject()在堆空间开辟的, 而不是像其他几个是基于地址偏移的开辟的栈空间.



### 三、new的变量在栈还是堆?



那么对于new出来的变量,是一定在heap中开辟的吗,我们来看看



```go
package main

func foo(arg_val int) (*int) {

    var foo_val1 * int = new(int);
    var foo_val2 * int = new(int);
    var foo_val3 * int = new(int);
    var foo_val4 * int = new(int);
    var foo_val5 * int = new(int);


    //此处循环是防止go编译器将foo优化成inline(内联函数)
    //如果是内联函数，main调用foo将是原地展开，所以foo_val1-5相当于main作用域的变量
    //即使foo_val3发生逃逸，地址与其他也是连续的
    for i := 0; i < 5; i++ {
        println(arg_val, foo_val1, foo_val2, foo_val3, foo_val4, foo_val5)
    }

    //返回foo_val3给main函数
    return foo_val3;
}


func main() {
    main_val := foo(666)

    println(*main_val, main_val)
}
```



我们将foo_val1-5全部用new的方式来开辟, 编译运行看结果



```bash
$ go run pro_3.go 
666 0xc000030728 0xc000030720 0xc00001a0e0 0xc000030738 0xc000030730
666 0xc000030728 0xc000030720 0xc00001a0e0 0xc000030738 0xc000030730
666 0xc000030728 0xc000030720 0xc00001a0e0 0xc000030738 0xc000030730
666 0xc000030728 0xc000030720 0xc00001a0e0 0xc000030738 0xc000030730
666 0xc000030728 0xc000030720 0xc00001a0e0 0xc000030738 0xc000030730
0 0xc00001a0e0
```



很明显, `foo_val3`的地址`0xc00001a0e0`依然与其他的不是连续的. 依然具备逃逸行为.



### 四、逃逸规则



我们其实都知道一个普遍的规则，就是如果变量需要使用堆空间，那么他就应该进行逃逸。但是实际上Golang并不仅仅把逃逸的规则如此泛泛。Golang会有很多场景具备出现逃逸的现象。



一般我们给一个引用类对象中的引用类成员进行赋值，可能出现逃逸现象。可以理解为访问一个引用对象实际上底层就是通过一个指针来间接的访问了，但如果再访问里面的引用成员就会有第二次间接访问，这样操作这部分对象的话，极大可能会出现逃逸的现象。



Go语言中的引用类型有func（函数类型），interface（接口类型），slice（切片类型），map（字典类型），channel（管道类型），*（指针类型）等。



那么我们下面的一些操作场景是产生逃逸的。



#### 逃逸范例一



`[]interface{}`数据类型，通过`[]`赋值必定会出现逃逸。



```go
package main

func main() {
    data := []interface{}{100, 200}
    data[0] = 100
}
```



我们通过编译看看逃逸结果



```bash
aceld:test ldb$ go tool compile -m 1.go

1.go:3:6: can inline main
1.go:4:23: []interface {}{...} does not escape
1.go:4:24: 100 does not escape
1.go:4:29: 200 does not escape
1.go:6:10: 100 escapes to heap
```



我们能看到，`data[0] = 100` 发生了逃逸现象。



#### 逃逸范例二



`map[string]interface{}`类型尝试通过赋值，必定会出现逃逸。



```go
package main

func main() {
    data := make(map[string]interface{})
    data["key"] = 200
}
```



我们通过编译看看逃逸结果



```go
aceld:test ldb$ go tool compile -m 2.go
2.go:3:6: can inline main
2.go:4:14: make(map[string]interface {}) does not escape
2.go:6:14: 200 escapes to heap
```



我们能看到，`data["key"] = 200` 发生了逃逸。



#### 逃逸范例三



`map[interface{}]interface{}`类型尝试通过赋值，会导致key和value的赋值，出现逃逸。



```go
package main

func main() {
    data := make(map[interface{}]interface{})
    data[100] = 200
}
```



我们通过编译看看逃逸结果



```go
aceld:test ldb$ go tool compile -m 3.go
3.go:3:6: can inline main
3.go:4:14: make(map[interface {}]interface {}) does not escape
3.go:6:6: 100 escapes to heap
3.go:6:12: 200 escapes to heap
```



我们能看到，`data[100] = 200` 中，100和200均发生了逃逸。



#### 逃逸范例四



`map[string][]string`数据类型，赋值会发生`[]string`发生逃逸。



```go
package main

func main() {
    data := make(map[string][]string)
    data["key"] = []string{"value"}
}
```



我们通过编译看看逃逸结果



```bash
aceld:test ldb$ go tool compile -m 4.go
4.go:3:6: can inline main
4.go:4:14: make(map[string][]string) does not escape
4.go:6:24: []string{...} escapes to heap
```



我们能看到，`[]string{...}`切片发生了逃逸。



#### 逃逸范例五



`[]*int`数据类型，赋值的右值会发生逃逸现象。



```go
package main

func main() {
    a := 10
    data := []*int{nil}
    data[0] = &a
}
```



我们通过编译看看逃逸结果



```bash
 aceld:test ldb$ go tool compile -m 5.go
5.go:3:6: can inline main
5.go:4:2: moved to heap: a
5.go:6:16: []*int{...} does not escape
```



其中 `moved to heap: a`，最终将变量a 移动到了堆上。



#### 逃逸范例六



`func(*int)`函数类型，进行函数赋值，会使传递的形参出现逃逸现象。



```go
package main

import "fmt"

func foo(a *int) {
    return
}

func main() {
    data := 10
    f := foo
    f(&data)
    fmt.Println(data)
}
```



我们通过编译看看逃逸结果



```bash
aceld:test ldb$ go tool compile -m 6.go
6.go:5:6: can inline foo
6.go:12:3: inlining call to foo
6.go:14:13: inlining call to fmt.Println
6.go:5:10: a does not escape
6.go:14:13: data escapes to heap
6.go:14:13: []interface {}{...} does not escape
:1: .this does not escape
```



我们会看到data已经被逃逸到堆上。



#### 逃逸范例七



- `func([]string)`: 函数类型，进行`[]string{"value"}`赋值，会使传递的参数出现逃逸现象。



```go
package main

import "fmt"

func foo(a []string) {
    return
}

func main() {
    s := []string{"aceld"}
    foo(s)
    fmt.Println(s)
}
```



我们通过编译看看逃逸结果



```bash
aceld:test ldb$ go tool compile -m 7.go
7.go:5:6: can inline foo
7.go:11:5: inlining call to foo
7.go:13:13: inlining call to fmt.Println
7.go:5:10: a does not escape
7.go:10:15: []string{...} escapes to heap
7.go:13:13: s escapes to heap
7.go:13:13: []interface {}{...} does not escape
 :1: .this does not escape
```



我们看到 `s escapes to heap`，s被逃逸到堆上。



#### 逃逸范例八



- `chan []string`数据类型，想当前channel中传输`[]string{"value"}`会发生逃逸现象。



```go
package main

func main() {
    ch := make(chan []string)

    s := []string{"aceld"}

    go func() {
        ch <- s
    }()
}
```



我们通过编译看看逃逸结果



```bash
aceld:test ldb$ go tool compile -m 8.go
8.go:8:5: can inline main.func1
8.go:6:15: []string{...} escapes to heap
8.go:8:5: func literal escapes to heap
```



我们看到`[]string{...} escapes to heap`, s被逃逸到堆上。



### 五、结论



Golang中一个函数内局部变量，不管是不是动态new出来的，它会被分配在堆还是栈，是由编译器做逃逸分析之后做出的决定。

- **栈分配**：适用于生命周期仅限于函数内部的局部变量，分配和释放速度快。
- **堆分配**：适用于生命周期超出函数作用域的变量，分配和释放速度较慢，且会增加垃圾回收的压力。

按理来说, 人家go的设计者明明就不希望开发者管这些,但是面试官就偏偏找这种问题问? 醉了也是.