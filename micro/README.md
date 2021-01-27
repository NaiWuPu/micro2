###复习开始
###micro 
```cmd
查看etcd里面的所有服务
micro --registry etcd --registry_address 8.136.142.17:2379 list services
查看服务源信息
micro --registry etcd --registry_address 8.136.142.17:2379 get service api.wuzi.com.test
呼叫呼叫
micro --registry etcd --registry_address 8.136.142.17:2379 call api.wuzi.com.test TestService.Call {\"id\":3}

micro call --address 127.0.0.1:8002 api.wuzi.com.test TestService.Call {\"id\":3}
```
### api调试失败 本地调用 8082端口调用失败了 端口丢失 未找到原因
```
go run main.go
post:http://127.0.0.1:8002
{"jsonrpc": "2.0", "method": "TestService.Call", "params": [{"id":3}], "id": 1}

apigw.bat
http://127.0.0.1:8003/test/TestService/call
{"id":3}
```