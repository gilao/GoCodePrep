以下是Git命令的详解，结合最新技术实践和网页内容整理：

---

### **一、基础命令**
1. **配置与初始化**  
   - `git config --global user.name/email`：设置全局用户名和邮箱  
   - `git init`：初始化本地仓库，生成`.git`目录  
   - `git clone <url>`：克隆远程仓库到本地，支持SSH/HTTPS协议  

2. **文件操作**  
   - `git add .`：将所有修改添加到暂存区  
   - `git commit -m "message"`：提交暂存区内容到本地仓库  
   - `git status`：查看工作区状态（修改/未跟踪文件）  
   - `git diff HEAD`：查看工作区与最新提交的差异  

3. **远程仓库交互**  
   - `git remote add origin <url>`：关联远程仓库  
   - `git push origin <branch>`：推送本地分支到远程  
   - `git pull origin <branch>`：拉取远程更新并合并  

---

### **二、分支管理**
1. **分支操作**  
   - `git branch <name>`：创建新分支  
   - `git checkout -b <name>`：创建并切换分支  
   - `git merge <branch>`：合并指定分支到当前分支  
   - `git branch -d <name>`：删除已合并的本地分支  

2. **分支查看**  
   - `git branch -a`：查看所有本地/远程分支  
   - `git log --graph --all`：图形化展示分支合并历史  

---

### **三、撤销与回退**
1. **提交回退**  
   - `git reset --soft HEAD^`：撤销提交但保留修改（暂存区恢复）  
   - `git reset --hard <commit>`：强制回退到指定版本（慎用）  
   - `git reflog`：恢复意外删除的提交或分支  

2. **文件恢复**  
   - `git checkout -- <file>`：撤销工作区未暂存的修改  
   - `git clean -f`：删除未跟踪的文件  

---

### **四、高级功能**
1. **历史追踪**  
   - `git bisect`：二分法定位引入问题的提交  
   - `git blame <file>`：查看文件每行修改的提交者  
   - `git log --grep="关键词"`：按提交信息关键词搜索  

2. **提交优化**  
   - `git commit --amend`：修改最近一次提交信息  
   - `git rebase -i HEAD~3`：交互式重写提交历史（需谨慎）  
   - `git stash`：临时保存未提交修改并切换分支  

---

### **五、协作与部署**
1. **标签管理**  
   - `git tag -a <tag> -m "message"`：创建带说明的附注标签  
   - `git push origin <tag>`：推送标签到远程  

2. **子模块**  
   - `git submodule add <url> <path>`：添加外部依赖库  
   - `git submodule update --init`：初始化子模块  

---

### **引用说明**
- 基础命令与配置：  
- 分支与合并：  
- 高级操作：  
- 标签与子模块：  
建议结合具体场景（如微服务多模块开发）进一步探讨分支策略或冲突解决技巧。