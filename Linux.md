## 文件操作
### 清空文件内容
命令：truncate -s 0 文件名<br>
效果：清空文件内容，支持大文件高效截断。

## 应用进程管理
### Supervisor
效果：Supervisor是一个进程管理工具，supervisord是它的守护进程，而supervisorctl是与supervisord进行交互的命令行工具。<br>

启动supervisord：在命令行中输入supervisord可以启动supervisord守护进程。<br>
使用supervisorctl连接到supervisord：在命令行中输入supervisorctl，即可连接到supervisord。<br>
常用命令：<br>
status：显示所有进程的状态信息，包括进程名称、进程ID、状态、启动时间等。<br>
start <process_name>：启动指定名称的进程。<br>
stop <process_name>：停止指定名称的进程。<br>
restart <process_name>：重启指定名称的进程。<br>
reload：重新加载配置文件，使之生效。<br>
update：根据更新的配置文件进行更新，并重启配置文件中更改的进程。<br>
add <config_file>：向supervisord添加一个新的进程配置。<br>
remove <process_name>：从supervisord中移除指定名称的进程。<br>
shutdown：关闭supervisord守护进程。<br>