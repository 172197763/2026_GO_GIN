# 2026_GO_GIN
gin练习项目-请勿关注！！！  
## 项目目录结构说明
项目名称  
├─ api ---gin框架业务逻辑<br>
├─ proto ---rpc服务接口定义<br>
└─ templates ---html模板<br>
.env.temp：系统配置信息<br>
## 项目初始化说明
1.初始化项目：go mod init 项目名称：gin_test<br>
2.安装依赖：go get -u github.com/gin-gonic/gin<br>
## 项目额外说明
go mod tidy<br>
>效果：一次性处理所有未引入的依赖<br>

air<br>
>安装：go install github.com/cosmtrek/air@latest<br>
>效果：可以监听文件变化并自动重启服务<br>
>提醒：需要配置文件：.air.toml<br>
