###使用了 micro的插件
```cmd
go get github.com/micro/go-plugins
```
### 生成代码 or gen.bat
```cmd
cd Models/protos
protoc --micro_out=../ --go_out=../ Prods.proto
```
### 一个三方插件 用来统一json ini 
```cmd
go get -u github.com/favadi/protoc-go-inject-tag
protoc-go-inject-tag -input=../Prods.pb.go
```
