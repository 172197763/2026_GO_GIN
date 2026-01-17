## 文件操作
### 清空文件内容
命令：truncate -s 0 文件名<br>
效果：清空文件内容，支持大文件高效截断。
### 创建多层目录
命令：mkdir -p 目录/{子目录1,子目录2……}<br>
效果：一次性创建多层目录

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

## mosquitto安装使用
### 安装
1. 创建目录：mkdir -p opt/mosquitto/{config,data,log}
2. 新增配置文件：config/mosquitto.conf
```
# mosquitto.conf
listener 1883
allow_anonymous false

# WebSocket
listener 9001
protocol websockets
allow_anonymous false


password_file /mosquitto/config/passwd

# 日志输出到 stdout（便于 docker logs 查看）
#log_dest stdout
log_dest file /mosquitto/log/mosquitto.log
log_type all

# 持久化（保留客户端会话、QoS 1/2 消息等）
persistence true
persistence_location /mosquitto/data/

```
3. 生成密码文件
```
# 在 mosquitto 目录下执行
docker run --rm -it \
  -v $(pwd)/config:/mosquitto/config \
  eclipse-mosquitto \
  mosquitto_passwd -c /mosquitto/config/passwd 用户名
# 然后输入两次密码即可
```
4. 编写Dockerfile
```
# mosquitto/docker-compose.yml
services:
  mosquitto:
    image: eclipse-mosquitto:latest
    restart: unless-stopped
    ports:
      - "1883:1883"   # MQTT TCP
      - "9001:9001"   # MQTT over WebSocket
    volumes:
      - ./config/mosquitto.conf:/mosquitto/config/mosquitto.conf
      - ./config/passwd:/mosquitto/config/passwd
      - ./data:/mosquitto/data
      - ./log:/mosquitto/log
    environment:
      - TZ=Asia/Shanghai

```
5. 启动服务<br>
在mosquitto目录下执行
```
# 启动服务
docker-compose up -d
# 停止服务
docker-compose down
# 查看日志
docker-compose logs -f mosquitto
#启动成功日志会显示
mosquitto  | 1705400000: Opening ipv4 listen socket on port 1883.
mosquitto  | 1705400000: Opening ipv6 listen socket on port 1883.
mosquitto  | 1705400000: Opening websockets listen socket on port 9001.
# 在yml文件目录下使用docker compose ps查看
mosquitto-mosquitto-1   eclipse-mosquitto:latest   "/docker-entrypoint.…"   mosquitto   47 minutes ago   Up 47 minutes   0.0.0.0:1883->1883/tcp, [::]:1883->1883/tcp, 0.0.0.0:9001->9001/tcp, [::]:9001->9001/tcp


```
6. 提示没有docker-compose 命令则需单独安装
```
# 下载最新版 compose 插件
sudo curl -SL https://github.com/docker/compose/releases/latest/download/docker-compose-linux-x86_64 -o /usr/local/bin/docker-compose

# 授权
sudo chmod +x /usr/local/bin/docker-compose

# 创建软链接（使 docker compose 可用）
sudo ln -s /usr/local/bin/docker-compose /usr/local/bin/docker-compose-plugin

# 验证安装
# 查看版本
docker compose version

# 或（旧命令兼容）
docker-compose --version
```