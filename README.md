## 升级说明

本分支主要对 charter10 章节内容作了升级，升级内容包括:

- 开发环境: go1.10 -> go1.14 (兼容 go1.10 或以上)
- 包版本管理: go modules
- go-micro: v1.1 -> v1.18 (主要对 consul 的注册逻辑进行了修改)
- 更新前端页面，更新样式，并支持分块上传

## 启动应用流程

- 配置数据库 (service/dbproxy/config/db.go)

```golang
MySQLSource = "test:test@tcp(127.0.0.1:3306)/fileserver?charset=utf8"
```

- 配置服务注册中心 consul (公共配置: config/service.go)

```golang
// RegistryConsul : 配置 consul
func RegistryConsul() registry.Registry {
	return consul.NewRegistry(
		// TODO ip需根据实际情况来修改
		registry.Addrs("192.168.2.244:8500"),
	)
}
```

- 配置下载服务 (service/download/config/config.go)

```golang
// DownloadEntry : 配置上传入口地址，用于 apigw 微服务动态获取、客户端获取后使用
var DownloadEntry = "127.0.0.1:38080"

// DownloadServiceHost : 上传服务监听的地址
var DownloadServiceHost = "0.0.0.0:38080"
```

- 配置上传服务 (service/upload/config/config.go)

```golang
// UploadEntry : 配置上传入口地址，用于 apigw 微服务动态获取、客户端获取后使用
var UploadEntry = "127.0.0.1:28080"

// UploadServiceHost : 上传服务监听的地址
var UploadServiceHost = "0.0.0.0:28080"
```

- 其他公共配置 ceph/oss/rabbitmq 等 (在 config 目录下配置)

- 批量启动微服务

```bash
# 注意脚本中的GOPATH / GOROOT 路径，要配置正确
./service/start-all.sh
```

- 批量关闭微服务

```bash
./service/stop-all.sh
```

- 前端网盘首页效果

<img src="/doc/home.png"></img>
