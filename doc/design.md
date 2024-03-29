## design

### overview

代码生成工具提供了 ddd，trpc等 crud代码 web框架模板

trpc模板采用trpc的目录规范实现，具体可以查看相关iwiki文档 [tRPC-Go目录规范](https://iwiki.woa.com/pages/viewpage.action?pageId=655870569)

ddd模板采用面向领域编程设计的思想实现的一套crud以及框架模板

### trpc目录规范

- 存在的问题

1. entity 实体类 可以理解，但为什么有错误码、远程配置等工具或者配置相关的东西呢

2. repo作为存储层（顾名思义），为什么有调用外部服务的逻辑呢，调用mysql跟调用其他的rpc服务，应该不是同一个层次的东西

这两点没得到满意的回答，包名存在很强的误导性

3. <b>ddd存在的问题，他都会存在，目录规范只是单纯设计了项目的root级结构，实际开发过程需要二次设计</b>

- 带来的好处

1. 共建合作提供统一规范

### ddd目录规范

```shell
.
├── ENGINEERING_GUIDE_GOLANG.md
├── README.md
├── conf // 存放配置相关 代码生成
│   ├── cache.go // 缓存配置相关 代码生成
│   └── rainbow_conf.go // 七彩石远程配置，代码生成
├── domain // ddd 域名模块，该目录有微服务领域里面的上下界，上下界本质是该微服务中的对象，例如权限控制微服务中，上下界就是user、role、permission等对象
│   ├── api // api 定义每个对象的接口定义 相当于 pb中的rpc定义 代码生成
│   │   └── user.go
│   ├── model // 对象model层，存储db的对象结构体，代码生成
│   │   └── user.go
│   ├── register// 对象服务注册 代码生成
│   │   └── user.go
│   └── user // 用户的crud主要逻辑 代码生成
│       ├── controller // 接口层的代码，实现了pb中rpc的接口 代码生成
│       │   ├── auth.go
│       │   ├── create.go
│       │   ├── delete.go
│       │   ├── get.go
│       │   ├── list.go
│       │   ├── task.go
│       │   └── update.go
│       ├── dao // 存储层的代码，支持cache 代码生成
│       │   ├── create.go
│       │   ├── dao.go
│       │   ├── delete.go
│       │   ├── get.go
│       │   ├── list.go
│       │   ├── save.go
│       │   └── update.go
│       └── server// server层代码，具体逻辑可以在这里实现 代码生成
│           ├── create.go
│           ├── delete.go
│           ├── get.go
│           ├── list.go
│           ├── server.go
│           └── update.go
├── go.mod
├── go.sum
├── inst // 单例，数据库实例 服务管理等 代码生成
│   ├── init.go
│   ├── inst.go
│   ├── mysql.go
│   ├── redis.go
│   └── server_manager.go
├── main.go
├── makefile // makefile ci工具
├── role.proto
├── stub
│   └── git.code.oa.com
│       └── trpcprotocol
│           └── tab
│               └── auth_center
│                   ├── go.mod
│                   ├── user.pb.go
│                   ├── user.proto
│                   ├── user.trpc.go
│                   └── user_service_mock.go
├── test
│   ├── Create.json
│   ├── Del.json
│   ├── Get.json
│   ├── List.json
│   └── Update.json
├── trpc_go.yaml
├── user_service.go
└── user_service_test.go
```



- 存在的问题

1. 过于面向对象，需要开发者有很好的抽象能力

   例如我们用关系型数据库 mysql 的时候，经常会设计一些关联表，例如User-Role表，按一张表一个对象的设计方式，这里会需要一个UserRoleRelation的对象，但严格上他是不对外暴露，只是关系型数据库表设计需要而存在的对象，这里的crud代码生成可能只需要dao层即可

​		另外在review 服务设计方案发现，开发者对对象定义也有自己的理解。好的对象设计，会有更好的模块化结构

- 带来的好处

1. 更好的易读性，cr友好

   根据ddd的思路，在一个业务域下存在一个上下界，例如权限控制这个域下，存在 User、Role、Permission、Token（optional）等基于rbac的对象，

   各个对象拥有的职能服务清晰明了，对于新增需求可以快速明确设计哪些对象的改动，更好的提供技术方案以及准确的工程排期

2. 更加结构化的层级，代码生成工具友好

   通过分层结构设计，抽离各个对象共通的逻辑，更好的整合项目，对代码复用，代码生成友好

3. 依赖其他域下的微服务更加简单

   ```go
   //ServiceManager ...
   type ServiceManager struct {
   	UserService api.UserService // 当前微服务对象，即同个域下的其他对象，本地代码调用，不需要走rpc
   	TaskService taskPB.TaskServiceService // 其他微服务对象，走rpc
   }
   ```

   通过 Inst.M.UserService 调用本进程的user对象的方法

   通过 Inst.M.TaskService 调用其他微服务的Task对象的方法

   可以看出，对于开发者，只需要用相同的调用方式，就可以获取到期望的对象服务，后续框架升级可以更好的屏蔽各个微服务的差别，让使用者感觉就是本地调用

4. 解决trpc目录规范存在的问题

   