# Code Generator

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