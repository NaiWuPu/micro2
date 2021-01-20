# micro2 AND etcd demo

## GET proto.go
```cmd
protoc --go_out=. --micro_out=. ./proto/hello.proto
```


## service
```
go run service.go
```

## Linux->etcd
```
micro --registry=etcd --registry_address=180.76.233.214:2379 list services
```

micro 2.0 goBin
```
链接：https://pan.baidu.com/s/1rvO_lHAVIVjwjjWEi7SzMA 
提取码：u5pu
```