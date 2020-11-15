# grpc-serverでのAPI開発を行う

# grpc-server
+ protoc --go_out=plugins=grpc:../pb cat.proto


# your-service

```
$ protoc -I . \
--go_out ./pb --go_opt paths=source_relative \
--go-grpc_out ./pb --go-grpc_opt paths=source_relative \
./proto/your_service.proto

```

# gateway作成

```
$ protoc -I./proto --grpc-gateway_out ./pb \
--grpc-gateway_opt logtostderr=true \
--grpc-gateway_opt paths=source_relative \
--grpc-gateway_opt generate_unbound_methods=true \
./proto/your_service.proto
```


```
$ protoc -I./proto --grpc-gateway_out -I$(GOPATH)/src \
--grpc-gateway_opt logtostderr=true \
--grpc-gateway_opt paths=source_relative \
--grpc-gateway_opt generate_unbound_methods=true \
./proto/your_service.proto


```