### 项目介绍
***
Crab 是一个基于Golang的应用脚手架,使用了热门库Gin,Gorm的组合可以帮助你快速构建一个高效、可靠的应用程序。

### 项目使用
***
在使用项目前需看modules目录下面例子在进行代码编写,推荐后端CRUD代码进行example方式编写极大提高效率。也可以按照自己喜欢的方式来modules设计极大提高了自主性！

### 优势
***
* **结构设计容易:** 刚刚学会GO语言就可以理解
* **内外分层:** 外层代码推荐写入pkg包在主程序引入即可,业务代码写在modules包
* **模块化设计:** 采用modules进行模块化设计,只用关心modules内部的代码
* **易升级:** Crab迭代不会影响modules与pkg(少量)业务代码,方便直接替换
* **代码简洁:** Crab使用了较多的 internal包进行同包名隔离,在编写代码时避免过多重复包名||超长包名||自定义包
* **协同开发:** 理论来说只需要定义好模型,modules下定义好包,即可像微服务一样多人开发

### 劣势
***
* **重复代码过多:** 因采用modules设计那么modules下包越多理论来说重复代码就会更多
* **server与orm的强依赖:** Crab没有做server与orm的代码解藕,如果库停止更新将是灾难。故选择最热门两个库

### 项目目录
```

├── cmd
│   └── internal
├── docs
├── model
│   └── example
├── modules
│   ├── example
│   │   └── internal
│   │       ├── controller
│   │       ├── middleware
│   │       ├── service
│   │       ├── util
│   │       └── vo
│   └── example2
│       └── internal
│           ├── controller
│           ├── middleware
│           ├── repo
│           │   └── impl
│           ├── service
│           │   └── impl
│           └── vo
└── pkg
    ├── cache
    ├── config
    ├── jwt
    ├── log
    ├── server
    └── sql

```