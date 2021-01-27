cd Services/protos
protoc --go_out=plugins=grpc:../../ServiceGW test.proto
protoc --grpc-gateway_out=logtostderr=true:../../ServiceGW test.proto
cd .. && cd ..