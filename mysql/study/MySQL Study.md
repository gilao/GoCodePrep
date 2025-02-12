# MySQL Study

# 1.数据库概述

## 1.1、数据库的好处

1. 将数据持久化到本地；
2. 提供结构化查询功能

## 1.2、数据库的常见概念

1、DB：数据库，存储数据的仓库

2、DBMS：数据库管理系统，又称为数据库软件或者数据库产品，用于创建和管理数据库，常见的有MySQL、Oracle、SQL Server

3、DBS：数据库系统，数据库系统是一个通称，包括数据库、数据库管理系统、数据库管理人员等，是最大的范畴

4、SQL：结构化查询语言，用于和数据库通信的语言，不是某个数据库软件特有的，而是几乎所有的主流数据库软件通用的语言

## 1.3、数据库的存储特点

1、数据存放到表中，然后表再放到库中

2、一个库中可以有多张表，每张表具有唯一的表名用来标识自己

3、表中有一个或多个列，列又称为“字段”，相当于Java中“属性”

4、表中的每一行数据，相当于Java中“对象”

## 1.4、数据库的常见分类

1、关系型数据库：MySQL、Oracle、DB2、SQL Server

2、非关系型数据库：

 a：键值存储数据库：Redis、Memcached、MemcacheDB

 b：列存储数据库：HBase、Cassandra

 c：面向文档的数据库：MongDB、CouchDB

 d：图形数据库：Neo4J

## 1.5、SQL语言的分类

1、DQL：数据查询语言：select、from、where

2、DML：数据操作语言：insert、update、delete

3、DDL：数据定义语言：create、alter、drop、truncate

4、DCL：数据控制语言：grant、revoke

5、TCL：事务控制语言：commit、rollback

# 第二章 MySQL概述

## 2.1、MySQL的背景

## 2.2、MySQL的优点

1、成本低、开源免费

2、性能高、移植性好

3、体积小、便于安装

## 2.3、MySQL的安装

## 2.4、MySQL的启动

## 2.5、MySQL的停止

## 2.6、MySQL的登录

## 2.7、MySQL的退出

# 第三章 DQL语言

## 3.1、基础查询

### 一、语法

```
SELECT 查询列表 FROM 表名;
```

### 二、特点

1. 查询列表可以是字段、常量、函数、表达式
2. 查询结果是一个虚拟表

### 三、示例

1. 查询单个字段

   ```sql
   SELECT 字段名 FROM 表名;
   ```

2. 查询多个字段

   ```
   SELECT 字段名,字段名 FROM 表名;
   ```

3. 查询所有字段

   ```
   SELECT * FROM 表名;
   ```

4. 查询常量

   ```
   SELECT 常量名;
   ```

5. 查询函数

   ```
   SELECT 函数名（实参列表）;
   ```

6. 查询表达式

   ```
   SELECT 100/25;
   ```

7. 起别名

   ```
   SELECT 字段名 AS "别名" FROM 表名;
   ```

   注意：别名可以使用单引号、双引号引起来，当只有一个单词时，可以省略引号，当有多个单词且有空格或特殊符号时，不能省略，AS可以省略

8. 去重复

   ```
   SELECT DISTINCT 字段名 FROM 表名;
   ```

9. 做加法

   ```
   SELECT 数值+数值;直接运算
   ```

   ```
   SELECT 字符+数值;首先先将字符转换为整数，如果转换成功，则继续运算，如果转化失败，则默认为0，然后继续运算；
   ```

   ```
   SELECT NULL+数值; NULL和任何数值参与运算结果都是NULL;
   ```

10. 【补充】ifnull 函数

    功能：判断某字段或表达式是否null；如果为null，返回指定的值，否则返回原本的值

    ```
    SELECT IFNULL(字段名，指定值) FROM 表名;
    ```

11. 【补充】isnull 函数

    功能：判断某字段或表达式是否为null，如果是null，则返回1，否则返回0；

    ```
    SELECT ISNULL(字段名) FROM 表名;
    ```

    

### 3.2、条件查询

#### 一、语法

```
SELECT 查询列表 FROM 表名 WHERE 筛选条件;
```

### 二、分类

1、条件运算符：>、>=、<、<=、=、<=>、!=、<>

2、逻辑运算符：and、or、not

3、模糊运算符：

 A、like：%任意多个字符、_任意单个字符，如果有特殊字符，需要使用escape转义

 B、between and

 C、not between and

 D、in

 E、is null

 F、is not null

### **三、演示**

1、查询工资>12000的员工信息

```SQL
SELECT 
*
FROM
employess
WHERE salary > 1200;-
```

