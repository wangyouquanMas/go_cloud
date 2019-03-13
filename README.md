# distributed-fileserver
基于golang实现的一种分布式云存储服务

- master分支: 通过原生net/http实现各接口功能
- gin分支: 通过Gin框架来改造(微服务化章节之后主要基于Gin框架来进行演示)

进度：
* [x] 简单的文件上传服务
* [x] mysql存储文件元数据
* [x] 账号系统, 注册/登录/查询用户或文件数据
* [x] 基于帐号的文件操作接口
* [x] 文件秒传功能
* [x] 文件分块上传/断点续传功能
* [x] 搭建及使用Ceph对象存储集群
* [x] 使用阿里云OSS对象存储服务
* [ ] 使用RabbitMQ实现异步任务队列
* [ ] 微服务化(API网关, 服务注册, RPC通讯)
* [ ] CI/CD(持续集成)

手记博客 [http://www.imooc.com/u/6198190](http://www.imooc.com/u/6198190)
