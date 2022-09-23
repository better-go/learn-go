# go-zero:

## 说明：

- 此示例 repo， 可以当作模板
- 修复了 goctl 生成的一些不好的点， rpc 示例， 可以当模板

```ruby
cd learn-go/

# pb 新增api， 动态生成新的模板代码：
task try:zero:gen:pb

```

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

## 参考：

> 文档：

- https://go-zero.dev/cn/docs/goctl/other
- https://go-zero.dev/cn/docs/goctl/api
- https://go-zero.dev/cn/docs/goctl/zrpc
- https://go-zero.dev/cn/docs/goctl/model

> 目录结构设计：

- https://go-zero.dev/cn/docs/advance/service-design

> 参考示例：

- https://go-zero.dev/cn/docs/eco/showcase
- https://github.com/Mikaelemmmm/go-zero-looklook
