### etcd 安装,随便找个安装方法
```cmd
https://github.com/etcd-io/etcd/releases
```
### etcd webui 安装 访问端口 2382
```cmd
docker run -it -d --name etcdkeeper \
-p 2382:8080 \
deltaprojects/etcdkeeper
```

###开始安装go micro
```
go get -u github.com/micro/go-micro
```
###未指定地址，使用命令行参数启动server (不推荐使用/少使用)
### 所以双击 prod.bat 就可以运行 多个server
```cmd
go run main.go --server_address :8001
```

###使用了 micro的插件
```cmd
go get github.com/micro/go-plugins
```


