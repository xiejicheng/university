### rpm 包
- rpm 包的格式
  - `vim-common`-`8.0.10.5`.`el7`.`x86_64`.rpm
- vim-common：软件名称
- 8.0.10.5：软件版本
- el7：系统版本
- x86_64：平台

#### rpm 命令
- -q 查询软件包
- -i 安装软件包
- -e 卸载软件包

### yum 命令常用选项
- install 安装软件包
- remove 卸载软件包
- list| grouplist 查看软件包
- update 升级软件包

#### yum依赖源改为国内镜像(ali)
1.备份
```
mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.backup
```
2.下载新的CentOS-Base.repo 到/etc/yum.repos.d/
```
# CentOS7
wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo
```

### 前台进程与后台进程
- `&` 写在命令后面表示设置此进程为后台进程
例如：
```
# 将a.sh文件在后台运行
$ ./a.sh &
```
一般默认情况下，我们在bash执行的命令(命令执行本质是开启一个进程)是前台进程，此时进程将shell占据了，我们无法进行其他的操作。而对于一些无需交互的进程，我们想让它们在后台运行，可以在启动命令的时候在命令后面加一个`&`实现此目的。

使用了`&`将一个进程放置到后台运行的时候，Bash会提示这个进程的进程ID。
> 每个进程都有唯一的进程ID，可以使用进程ID来暂停、恢复、终止相应的进程。

- `jobs` 命令可以显示当前终端正在运行的进程，包括前台运行和后台运行的进程。它对每个正在执行中的进程分配一个序号，注意的是此序号不是进程ID，此序号可以进行多种操作。

- `fg` 命令可以将后台运行的进程放到前台运行。通过`jobs` 命令提供的进程序号，再在序号前面加上 `%`符号，就可以把相应的进程任务放到前台运行。
如果进程任务是暂停状态，`fg`命令会将它启动起来。

- 使用 `Ctrl+z` 组合键可以将前台运行的任务暂停。
- `bg` 命令会将任务放置到后台运行，如果任务是暂停状态的，会被启动。

想在在另一个终端终止某个进程，可以使用 `kill` 命令：
```
$ kill -s TERM <PID>
# PID 为进程ID
```
- `TERM` ：终止一个进程
- `STOP` 等同于 Ctrl+z 组合键的效果，使进程暂停
- `CONT`：将暂停的进程启动起来
- `KILL`：强制终止进程

### 进程优先级调整
#### 调整优先级
- `nice` 范围从 -20 到 19， 值越小优先级越高，抢占资源越多。
- `renice` 重新设置优先级
```
$ nice -n -20 vi &  //设置优先级为 -20
$ renice -n 15 8924 //将进程ID为8924的任务的优先级调整为15
```

> ps、top 命令中的`NI`就是优先级

### 进程间通信
#### 信号是进程间通信方式之一
- Linux kill命令用于删除执行中的程序或工作。
```
# kill -l
1) SIGHUP     2) SIGINT     3) SIGQUIT     4) SIGILL     5) SIGTRAP
6) SIGABRT     7) SIGBUS     8) SIGFPE     9) SIGKILL    10) SIGUSR1
11) SIGSEGV    12) SIGUSR2    13) SIGPIPE    14) SIGALRM    15) SIGTERM
16) SIGSTKFLT    17) SIGCHLD    18) SIGCONT    19) SIGSTOP    20) SIGTSTP
21) SIGTTIN    22) SIGTTOU    23) SIGURG    24) SIGXCPU    25) SIGXFSZ
26) SIGVTALRM    27) SIGPROF    28) SIGWINCH    29) SIGIO    30) SIGPWR
31) SIGSYS    34) SIGRTMIN    35) SIGRTMIN+1    36) SIGRTMIN+2    37) SIGRTMIN+3
38) SIGRTMIN+4    39) SIGRTMIN+5    40) SIGRTMIN+6    41) SIGRTMIN+7    42) SIGRTMIN+8
43) SIGRTMIN+9    44) SIGRTMIN+10    45) SIGRTMIN+11    46) SIGRTMIN+12    47) SIGRTMIN+13
48) SIGRTMIN+14    49) SIGRTMIN+15    50) SIGRTMAX-14    51) SIGRTMAX-13    52) SIGRTMAX-12
53) SIGRTMAX-11    54) SIGRTMAX-10    55) SIGRTMAX-9    56) SIGRTMAX-8    57) SIGRTMAX-7
58) SIGRTMAX-6    59) SIGRTMAX-5    60) SIGRTMAX-4    61) SIGRTMAX-3    62) SIGRTMAX-2
63) SIGRTMAX-1    64) SIGRTMAX
```

### 守护进程
守护进程(daemon)是一类在后台运行的特殊进程，用于执行特定的系统任务。很多守护进程在系统引导的时候启动，并且一直运行直到系统关闭。另一些只在需要的时候才启动，完成任务后就自动结束。

`nohup`命令：如果你正在运行一个进程，而且你觉得在退出帐户时该进程还不会结束，那么可以使用nohup命令。该命令可以在你退出帐户/关闭终端之后继续运行相应的进程。
- `nohup`命令使得进程忽略hangup挂起信号

### screen 命令
screen命令用于多重视窗管理程序。
```
$ yum install screen // 安装
$ screen  // 进入screen环境
$ screen -ls  // 查看 screen 的会话
$ screen - r sessionid // 恢复会话
```
- `ctrl+a  d` 退出 srceen 环境

### Linux启动级别
- 0:关机  
- 1:单用户  
- 2:多用户状态无网络服务  
- 3:多用户状态有网络服务  
- 4:系统未使用保留给用户  
- 5:图形界面  
- 6:系统重启

### 如何将一个在前台已经运行的进程转到后台运行呢？
- ctrl + z 将进程变为 `Stopped` 状态并放到后台
- 可以通过 `jobs` 命令查看到在后台呈 `Stopped`状态的进程（前头的第一个数字为编号，我这里为1）
- 通过命令`bg %1` 使得进程在后台运行起来
- 再次通过命令 `jobs` 可以看到这个后台进程的状态为 `Running`