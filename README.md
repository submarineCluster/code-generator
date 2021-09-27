# Code Generator

[tab微服务框架使用文档](./doc/tab_micro_server.md)

[为什么这么设计](./doc/design.md)

## Overview

> code-generator 以资源为粒度，生成该资源的 API 代码，如常见的 CRUD
>
> - create
> - get
> - list
> - update
> - delete

## Resource Layer

### Store Plugin

> 通过实现store 的插件接口，实现 dao 层存储无关

### Model

> 资源的属性

### Dao 

> 资源的存储操作代码

### Controller

> 资源的 controller 层代码，用户可以将主要业务逻辑放在这

#### Protocol

>  资源协议存储，例如 pb 协议，thrift 协议等

### Test

#### PostMan

#### CURL

## Feature

- [x] CRUD code generator 
- [ ] Protocol File generator

## TODO

- standard Fields
- HTTP Server Gateway[option]
- cache layer
- page+snap
- mongo template
- test code: eg curl\postman\swagger
- mysql ceate-table utils
- standard lib: eg page

## Start

> go get -u git.code.oa.com/tencent_abtest/code-generator

### Example

#### generate crud code

> code-generator resource -n demo

该命令会在的当前gomod项目下的domain目录生成demo的crud 代码

具体参数可以查看help

> $ code-generator resource -h
> generate crud code for resource:
>
> code-generator resource will generator crud code for resource
>
> Usage:
>   code-generator resource [flags]
>
> Examples:
> code-generator resource -n codeGenerator
>
> Flags:
>       --apiServer            generator api server
>       --daoMetrics           report metrics for dao op (default true)
>   -h, --help                 help for resource
>   -n, --name string          name of resource
>   -s, --storageType string   storage-type, eg: mysql/mongo, mysql default (default "mysql")
>   -v, --verbose              verbose



## License

Copyright © 2020 ihago [git.yy.com]

# code-generator

一行简单的描述。(短描述)

## 项目介绍     README

本组织/项目的整体说明概览

## 快速上手     Getting Started

使用者如何快速上手使用本组织/项目，比如API列表和运维信息等。
https://iwiki.woa.com/pages/viewpage.action?pageId=978820183

### API

如果是对外提供服务API或公共库API，那么在这里**详细**描述你的接口，并且给出一个示例。

### 运维

这个章节需要描述当你服务出问题时的一些操作，比如：

- 服务降级
- 问题定位
- 监控查看
- 日志查看

## 常见问题     FAQ

本组织/项目的常见通用问题和官方解答

## 行为准则    Code Of Conduct

本组织/项目在代码协作方面需遵循的责任、范围、软件许可证、冲突解决等章程

## 如何加入    How To Join

本组织/项目有明确的如何加入和贡献的文字说明

## 团队介绍    Members

本组织/项目的角色分工、人名和联络方式、官方交流/沟通渠道

## 参考    References

假如你的项目参考或引用了其他项目资料，请列出来。

- [视频后台研发手册](https://git.code.oa.com/videobase/videonavi)

