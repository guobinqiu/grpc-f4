# gRPC Demo

注意都是后端对后端

## 安装

> https://github.com/guobinqiu/grpc-exercise/tree/master/go-as-serverside#go-as-server-side

## 4种模式

### 1. 请求-响应 (simple RPC)

替代 http api

### 2. 客户端推流 (client-side streaming RPC)

文件上传、数据上报

### 3. 服务端推流 (server-side streaming RPC)

文件下载、数据下发 

如果客户端是`浏览器`用`SSE`代替

### 4. 双向推流 (bidirectional streaming RPC 全双工)

聊天、音视频通话

如果客户端是`浏览器`用[websocket](https://github.com/guobinqiu/vue2-go-websocket-protobuf-demo)代替
