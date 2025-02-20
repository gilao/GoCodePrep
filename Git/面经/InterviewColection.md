### **一、基础概念类**
1. **Git是什么？它与SVN等版本控制系统有何区别？**  
   - Git是分布式版本控制系统，每个开发者本地拥有完整仓库副本，支持离线工作；SVN是集中式，依赖中央服务器。  
   - 核心差异：分布式架构、分支管理、数据完整性（SHA-1哈希）。

2. **解释Git的三个工作区域：工作区、暂存区、仓库**  
   - 工作区：编辑代码的地方；暂存区（Index）：暂存待提交修改；仓库：存储完整历史记录。

3. **如何创建并切换到新分支？**  
   - `git checkout -b <branch-name>` 或 `git switch -c <branch-name>`（Git 2.23+）。

---

### **二、实际操作类**
4. **合并分支时发生冲突，如何解决？**  
   - 步骤：`git status` 查看冲突文件 → 手动编辑冲突部分 → `git add <file>` 标记解决 → `git commit` 完成合并。

5. **如何撤销最后一次提交？**  
   - 未推送到远程：`git reset HEAD^`（保留工作区修改）或 `git commit --amend`（修改提交信息）。  
   - 已推送远程：`git revert <commit-hash>` 创建新撤销提交并推送。

6. **如何删除已合并的远程分支？**  
   - `git push origin --delete <branch-name>`。

---

### **三、进阶机制类**
7. **Git垃圾收集器（GC）的作用是什么？**  
   - 清理不再引用的对象（如废弃分支、大文件），释放存储空间，可通过 `git gc` 手动触发。

8. **解释Git的Fast-Forward合并与递归合并策略**  
   - Fast-Forward：直接移动分支指针（无新提交时）；递归合并：创建新合并提交（保留分支历史）。

9. **如何查找特定提交中修改的文件列表？**  
   - `git diff-tree -r --name-only <commit-hash>`。

---

### **四、协作与部署类**
10. **Git Flow工作流的核心流程是什么？**  
    - 主分支（master/main）用于发布，开发在feature分支，通过pull request合并，支持版本标签管理。

11. **如何安全地重置远程分支的提交历史？**  
    - 强制推送：`git push -f`，但需确保团队成员同步清理本地缓存。

12. **Git子模块（Submodule）的使用场景与命令**  
    - 场景：管理外部依赖库；命令：`git submodule add <url> <path>`、`git submodule update --init`。

---

### **五、原理与调试类**
13. **Git如何保证数据完整性？**  
    - 所有操作基于SHA-1哈希校验，任何文件或提交内容变化都会导致哈希变化。

14. **如何调试Git操作？**  
    - 使用 `git reflog` 查看分支操作历史，`git bisect` 二分法定位引入问题的提交。

---

### **六、高频陷阱类**
15. **`git pull` 和 `git fetch` 的区别？**  
    - `git fetch` 仅下载远程更新，需手动合并；`git pull` = `fetch + merge`。

16. **误删本地分支后如何恢复？**  
    - 通过 `git reflog` 找到分支的最后一个提交，重新创建分支。
