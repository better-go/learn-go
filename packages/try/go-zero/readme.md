# go-zero:

> 文档：

- https://go-zero.dev/cn/docs/goctl/other
- https://go-zero.dev/cn/docs/goctl/api
- https://go-zero.dev/cn/docs/goctl/zrpc
- https://go-zero.dev/cn/docs/goctl/model

## 示例服务：

- run:

```ruby

cd learn-go/

# go mod tidy:
task try:zero:tidy 

# run api:
task try:zero:run:api

# run rpc:
task try:zero:run:rpc

```

- api 验证：


```ruby

cd learn-go/

# api test:
task try:zero:api:test

```

### api 示例：

- new：

```ruby

goctl api new hello

```

### rpc 示例：

- new：

```ruby

goctl rpc new rpc
```

- 依赖服务发现机制， 默认基于 etcd 注册 / 发现服务。

