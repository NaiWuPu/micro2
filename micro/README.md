###复习开始
###micro 
```cmd
查看etcd里面的所有服务
micro --registry etcd --registry_address 8.136.142.17:2379 list services
查看服务源信息
micro --registry etcd --registry_address 8.136.142.17:2379 get service test
呼叫呼叫
micro --registry etcd --registry_address 8.136.142.17:2379 call test TestService.Call {\"id\":3}
```