# 项目要求
```go
扫描调度框架，建议实现如下功能:
1、扫描器一键发布（通过提供ssh地址、账号、密码，实现一键发布）
2、任务增、删、改、查，状态监控
3、扫描结果回传、展示（在调度框架上实现初步展示）
4、扫描结果与网络空间测绘平台对接，清洗、入库
```
# 一 整体设计

![](./screenshot/manager.png)

## 1 通讯逻辑
```go
通讯方式改为了中心节点向worker发送http心跳请求
work节点就是一个http的服务
中心调度系统确实也是可以分布式的
```

## 2 调度逻辑

中心节点遍历任务，抢锁后通过http请求分发任务

# 二 源码相关

## 1 静态资源的打包

所用的包
```go
Go 语言打包静态文件以及如何与Gin一起使用Go-bindata
```

安装方法
```go

go get -u github.com/jteeuwen/go-bindata/...

安装前需要先关闭go mod
go env -w GO111MODULE=off

#通过命令行加入GOBIN的PATH
export PATH=$PATH:$GOPATH/bin
#编辑启动配置文件，开机后自动加载这个路径
nano ~/.bashrc
#编辑完成后，重新加载环境变量到内存
source ~/.bashrc

```

前端修改后打包进go项目
```go
go-bindata -o=core/utils/asset/asset.go -pkg=asset web/crocodile/... sql
```

# 三 项目的安装

数据库是先要新建的   表和数据是通过代码初始化的


#
vscan -host {ip} -p {port} -json -o {res}