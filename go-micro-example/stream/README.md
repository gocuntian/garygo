# Stream

一个流服务和两个客户端，一个流式rpc客户端和一个使用websockets的客户机的例子。

## Contents

- server - is the service
- client - is the rpc client
- web - is the websocket client

## Prereqs

微服务需要一个发现系统，以便他们可以找到对方。 Micro默认使用领事，但是它很容易swapped out with etcd，kubernetes或其他各种系统交换出来。
Install consul
```shell
brew install consul
```
Run Consul

```shell
consul agent -dev
```
## Proto
 #protoc -I proto/  proto/stream.proto --go_out=plugins=micro:proto

## Run the example

Run the service

```shell
go run server/main.go
```

Run the client

```shell
go run client/main.go
```

Run the micro web reverse proxy for the websocket client

``` shell
micro web
```

Run the websocket client

```shell
cd web # must be in the web directory to serve static files.
go run main.go
```

Visit http://localhost:8082/stream and send a request!

And that's all there is to it.
