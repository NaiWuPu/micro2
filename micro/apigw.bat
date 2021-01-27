set MICRO_REGISTRY=etcd
set MICRO_REGISTRY_ADDRESS=8.136.142.17:2379
micro api --address 127.0.0.1:8003 --namespace=api.wuzi.com --handler=api --resolver=host
