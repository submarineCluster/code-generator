# tab-micro-server

代码生成工具篇

可以查看 工具helper

```shell
✘ shaohui@LEOSHLI-MB1  ~/Code/tencent/tencent-abtest  code-generator resource -h
generate crud code for resource:

code-generator resource will generator crud code for resource

Usage:
  code-generator resource [flags]

Examples:
code-generator resource -n codeGenerator

Flags:
      --apiServer            API server代码生成，TODO 暂不需要
  -a, --app string           appName default TAB, 生成协议文件时需要传入 (default "TAB")
  -c, --cacheEnable          开启对象数据redis缓存，一般不需要，除非像用户等经常需要Get获取的
      --daoMetrics           指标上报代码生成，TODO (default true)
  -h, --help                 help for resource
  -n, --name string          资源对象名称，必填，如demo
  -p, --protoOnly            单纯生成proto文件
      --server string        serverName, 一般用123上的服务名，请使用小蛇式如 mab_scheduler
  -s, --storageType string   storage-type, eg: mysql/mongo, mongo default， TODO (default "mongo")
  -t, --templateDir string   path of template-code，推荐使用ddd，备选trpc, 例如 code-generator resource -n demo -t ddd
  -v, --verbose              verbose

Global Flags:
      --config string   config file (default is $HOME/.code-generator.yaml)
```

## 使用方式

### 生成协议

选择进入代码目录，会在该目录下生成对象的proto协议文件，以权限管理服务为例，生成user对象协议

```shell
code-generator resource -n user -t ddd -p --server auth_center
```

当前目录可以看到user.proto文件

```shell
 shaohui@LEOSHLI-MB1  ~/Code/tencent/tencent-abtest  ll | grep user
-rw-r--r--   1 shaohui  staff   2.6K  7  5 21:42 user.proto
```

### 使用vepc工具生成项目

```shell
vepc create --lang=go --protofile=./user.proto --git-path=tencent_abtest/auth-center --bk-project=tab --outputdir=./auth-center
```

会在当前目录生成auth-center的项目，如果vepc生成代码遇到问题，可以多次允许，直到稳定报错或者成功（vepc工具问题

### 生成对象crud

进入auth-center，通过以下命令生成user的增删改查代码

这里的crud 项目模板推荐使用ddd，如果要用trpc-go编写的的令人发指的规范（logic、repo、entity）的话，使用过程有问题，工具维护者没眼看

```shell
code-generator resource -n user -t ddd --server auth_center
```

```shell
 shaohui@LEOSHLI-MB1  ~/Code/tencent/tencent-abtest/auth-center   master ±✚  tree    
.
├── ENGINEERING_GUIDE_GOLANG.md
├── README.md
├── auth-center
├── conf
│   ├── cache.go
│   └── rainbow_conf.go
├── domain
│   ├── api
│   │   └── user.go
│   ├── model
│   │   └── user.go
│   ├── register
│   │   └── user.go
│   └── user
│       ├── controller
│       │   ├── auth.go
│       │   ├── create.go
│       │   ├── delete.go
│       │   ├── get.go
│       │   ├── list.go
│       │   ├── task.go
│       │   └── update.go
│       ├── dao
│       │   ├── create.go
│       │   ├── dao.go
│       │   ├── delete.go
│       │   ├── get.go
│       │   ├── list.go
│       │   ├── save.go
│       │   └── update.go
│       └── server
│           ├── create.go
│           ├── delete.go
│           ├── get.go
│           ├── list.go
│           ├── server.go
│           └── update.go
├── go.mod
├── go.sum
├── inst
│   ├── init.go
│   ├── inst.go
│   ├── mysql.go
│   ├── redis.go
│   └── server_manager.go
├── main.go
├── makefile
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

生成后，需要稍做修改，将main方法中的

```go
pb.RegisterUserServiceService(s, &userServiceServiceImpl{})
```

改为生成的对象，并删除框架生成的对象&userServiceServiceImpl{} 以及 对应的rpc方法

```go
pb.RegisterUserServiceService(s, &controller.UserController{Server: server.NewUserServer(dao.NewStore(common_dao.NewDao(inst.DBClient)))})
```

### 检查配置

查看conf目录下的rainbow_conf.go 文件配置

```go
const (
	serverConfName   = "/*替换成自己的配置文件 推荐serverConfig*/"
	rainbowURL       = "http://api.rainbow.oa.com:8080"       // 七彩石的host
	rainbowAppID     = "55a109c7-4a17-4658-84f4-e29812dfbe5b" // 七彩石的项目ID
	rainbowGroupName = "auth_center"                          // 七彩石的配置空间名，默认用serverName
)
```

### 编写对象属性

编辑domain/model/user.go

```go
//UserSpec ...
type UserSpec struct{}

//TableName ...
func (u *User) TableName() string {
	return "t_user"
}
```

在UserSpec中增加user对象的属性，例如BusinessCode、RoleCodeList，其他基本属性如ID，Name，工具已经生成，具体可以看User对象中的ObjectMeta

```go
//UserSpec ...
type UserSpec struct{
	BusinessCode string `json:"businessCode"`
	RoleCodeList []string `json:"roleList"`
}
```

```go
//User ...
type User struct {
	meta.ObjectMeta
	Spec *UserSpec       `gorm:"embedded"`
	Kind meta.ObjectKind `gorm:"-"`
}
```

### 创建db表结构

简单一个user例子

```mysql
CREATE TABLE `t_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL,
  `title` varchar(255) NOT NULL,
  `namespace` varchar(512) NOT NULL,
  `labels` varchar(512) NOT NULL,
  `annotations` varchar(512) DEFAULT NULL,
  `business_code` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 11 DEFAULT CHARSET = utf8mb4 COMMENT = 'tab平台用户模型'
```

### done，愉快写domain/user/server下的代码

权限管理也有很多个对象 除了user，还有role，permission、token，

同样用

```shell
code-generator resource -n role -t ddd --server auth_center
```

```shell
code-generator resource -n permission -t ddd --server auth_center
```

```shell
code-generator resource -n token -t ddd --server auth_center
```

