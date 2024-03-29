# learn-go

- go 第三方 lib 调研 / 验证.

> related:

- ✅ https://github.com/better-py/learn-py
- ✅ https://github.com/better-rs/learn-rs
- ✅ https://github.com/better-dart/learn-dart
- ✅ https://github.com/better-zig/learn-zig

## Projects：

- ✅️ monorepo: [packages](./packages/)

[//]: # (y23m01p01-xxx 项目编号规则)

| Project Name                           | Desc                    | Rank |  
|:---------------------------------------|:------------------------|:-----| 
| ✅️ [go-101](./packages/go-101)         | go 基础示例                 | ⭐️⭐️ |
| ✅️ [go-zero](packages/try/go-zero)     | go-zero 测试 demo， 功能验证   | xxx  |
| ✅️ [go-kratos](packages/try/go-kratos) | go-kratos 测试 demo， 功能验证 | xxx  |
| ✅️ [xxx]()                             | xxx                     | xxx  |
| ✅️ [xxx]()                             | xxx                     | xxx  |
| ✅️ [xxx]()                             | xxx                     | xxx  |
| ✅️ [xxx]()                             | xxx                     | xxx  |
| ✅️ [xxx]()                             | xxx                     | xxx  |


## Structures:

```ruby

➤ tree . -L 5
.
├── Taskfile.yml                             // 服务启动脚本
├── go.work
├── packages
│   └── try
│       ├── go-kratos                  // go-kratos 示例
│       │   ├── Taskfile.yml
│       │   ├── go.mod
│       │   └── main.go
│       └── go-zero                    // go-zero 最佳实践， 包含目录规范
│           ├── Taskfile.yml
│           ├── api                   // api 服务代码（自动生成）
│           │   └── main.go
│           ├── dart
│           ├── go.mod
│           ├── go.sum
│           ├── proto                 // go-kratos 服务所有定义文件， 统一管理
│           │   ├── api         // api 服务定义， 用于生成 api/ 服务目录代码
│           │   ├── model       // model 代码（自动生成）， 基于 sql 生成
│           │   ├── rpc         // rpc 服务定义， 用于生成 rpc/ 服务目录代码
│           │   └── sql         // sql 文件， 用于生成 proto/model/ CRUD 代码
│           └── rpc                   // rpc 服务代码（自动生成）
│               ├── main.go
└── tmp
    └── demo
        └── go.mod

23 directories, 17 files


```

## Requirements:

- ✅ go, 使用 workspace 模式
- ✅ go-task

```ruby
➤ go version
go version go1.18 darwin/arm64

go version go1.20.3 darwin/arm64

```