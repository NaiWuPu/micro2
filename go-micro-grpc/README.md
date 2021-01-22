### demo讲解
```cmd
通过访问
http://127.0.0.1:8001/v1/prods?size=20
{"size":20}
接口
将数据发送到 client 端 gin_main.go
gin_main.go 通过etcd 获取 prodService.client 的服务
通过绑定的接口，从 main.go 获取返回值
```
### 将wrapper 玩起来，中间件的东东
```cmd
https://github.com/micro/go-plugins/tree/master/wrapper
```
### 超时熔断机制 go-plugin 包括了此库
```cmd
https://github.com/afex/hystrix-go
```
### 使用方式
```cmd
server 端：go run main
client端：go run gin_main.go
```