### 安装 Git

- windows：git-scm.com 下载安装包
- linux/Mac：直接终端安装

#### 查看版本

```bash
$ git --version
```



### Git 配置

```bash
git config --global user.name "xxxxx"
git config --global user.email "xxxx@gmail.com"

```

#### 查看配置

```bash
$ cat ~/.gitconfig 
[user]
	name = xxxxx
	email = xxxxxd@gmail.com
```



### Git 基本工作流

#### 初始化项目

```bash
$ git init
$ ls -a
 .git
```

.git 为 git 的配置文件

`git add`  命令将文件内容添加到索引(将修改添加到暂存区)。也就是将要提交的文件的信息添加到索引库中。

```bash
# 将所有文件添加到暂存区
$ git add . 
```

`git commit`  命令用于将更改记录(提交)到存储库。将索引的当前内容与描述更改的用户和日志消息一起存储在新的提交中。(真正的将文件提交到版本控制中)

```bash
$ git commit -m 'describe...'
```

`git log` 命令用于显示提交日志信息。(查看 commit 的信息)

```bash
$ git log
```

`git status` 命令用于显示工作目录和暂存区的状态。使用此命令能看到那些修改被暂存到了, 哪些没有。

```bash
$ git status
```



### 修正错误的 commit

```bash
# 增补提交，会使用与当前提交节点相同的父节点进行一次新的提交，旧的提交将会被取消。
$ git commit --amend 
```



### 忽略不必要的文件

- `.gitignore`  这个文件的作用就是告诉Git哪些文件不需要添加到版本管理中。该文件一般放在根目录。

  只需要在.git同目录下新增.gitignore文件，然后添加不需要上次的目录即可，比如：

  ```bash
  /doc/   #过滤整个文件夹
  /target/
  /.idea/
  /offer.iml  #过滤文件
  
  #需要注意的是，gitignore还可以指定要将哪些文件添加到版本管理中：
  !*.zip    #开头多了一个感叹号
  !/mtk/one.txt
  
  #为什么要有两种规则呢？想象一个场景：我们只需要管理/mtk/目录中的one.txt文件，这个目录中的其他文件都不需要管理。那么我们就需要使用：
  
  /mtk/
  !/mtk/one.txt
  ```

#### 清除已经上传的多余文件

```bash
如果你添加.gitignore的时候，git里面已经上传了很多不需要的文件，则使用下面两个命令干掉他们
如果是文件夹：git rm -r --cached 文件夹名
如果是文件：git rm --cached 文件名
```



### Git 分支理解与应用

`git checkout` 命令用于切换分支或恢复工作树文件。

`git branch ` 命令用于列出，创建或删除分支。

`git merge` 命令用于将两个或两个以上的开发历史加入(合并)一起



### 处理合并冲突 Git merge conflict

 `git diff` 命令用于显示提交和工作树等之间的更改。此命令比较的是工作目录中当前文件和暂存区域快照之间的差异,也就是修改之后还没有暂存起来的变化内容。



### 设置 Git alias (快捷方式)

如果敲 git s 就表示 git status 那简单多了很多

我们只需要敲一行命令，告诉 Git，以后 s  就表示 status：

```bash
$ git config --global alias.s status
```

配置记录在 .git/.gitconfig 文件中的 [alias]



### Git stash 改善工作流

```bash
# 创建新分支
$ git checkout -b feature/new-design-2
# 在新分支中新新功能的代码
...
#此时，需要紧急修复一个 bug
#执行 git stash 将新功能代码放在一边
$ git stash
#返回主分支 master 修复 bug
$ git checkout master
...
#修复完 bug ，回到新分支
$ git checkout feature/new-design-2
#执行 pop 命令，将新功能代码恢复回来
$ git stash pop
...
#写完新功能后，回到 master 分支，使用 merge 合并分支
$ git checkout master
$ git merge feature/new-design-2
...
```



### GitHub 管理项目

- Issues：用来追踪各种想法,增强功能,任务,bug等。



### 理解 Git rebase

`git rebase` 命令在另一个分支基础之上重新应用，用于把一个分支的修改合并到当前分支。



### GitHub 贡献代码

- Fork
- git remote 
- 创建新分支
- pull 主分支
- 合并分支
- Pull requests

```bash
$ git remote add upstream
$ git pull upstream master #拿到最新代码
$ git checkout -b  #创建新分支
$ git commit #写完代码后 commit
$ git pull upstream master #获得最新代码
$ git checkout && git rebase master
$ git git push origin branch
```



### Github 使用小技巧

:fire: Trending 

github 邮件推送：github.com/explore/subscribe

#### 快捷键

- <>Code 面板下 ，按 t 键进入搜索模式，可以搜索我们想找的文件
- gi：进入 Issues 面板
- gp：进入 Pull requests 面板
- gc：进入 <>Code 面板
- s：打开搜索框

**Shift + ?  弹出所有快捷键操作方法**

