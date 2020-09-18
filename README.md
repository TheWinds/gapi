# GAPI

## 介绍

GAPI是一款实现了HTTP和GRPC之间相互转码的HTTP框架，GAPI从protobuf定义中生成HTTP API并完成了中间的调用过程。不同于[grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)，GAPI选择通过自己维护的元信息动态生成路由，这就达到了变更service protobuf定义后无需重启HTTP服务的效果，使基于GAPI实现动态网关成为可能。

## 特性

- [x] 动态生成 HTTP handler
- [x] 支持自定义请求中间件
- [x] 支持从 header / URL params / form / context 绑定参数
- [x] 支持返回字段的别名&默认值选项
- [x] 开箱即用的 httpjson callhander
- [x] 可选的 HTTP 透传

## 使用

### 安装
```sh
go get github.com/zhiduoke/gapi
```

### 使用

#### API定义
1. 使用protocol buffer 定义 gRPC 服务
```protobuf
syntax = "proto3";

package demo;

option go_package = ".;api";

message GreetRequest{
    int64 userId = 1;
    string name = 2;
}

message GreetReply{
    string msg = 2;
}

service DemoAPI {
    rpc Greet (GreetRequest) returns (GreetReply) {
    }
}
```

2. 增加GAPI注解

这一步增加了一些用于定义api的注解，具体含义在后面说明。

```diff
syntax = "proto3";

package demo;

option go_package = ".;api";

+
+ import "gapi/proto/annotation.proto";
+

message GreetRequest{
-   int64 userId = 1;
+   int64 userId = 1 [(gapi.from_context) = true, (gapi.alias) = "uid"];
    string name = 2;
}

message GreetReply{
    string msg = 2;
}

service DemoAPI {
+   option (gapi.server) = "service.demo";
+   option (gapi.default_handler) = "httpjson";
+   option (gapi.default_timeout) = 5000;
+   option (gapi.path_prefix) = "/v1/demo";
    rpc Greet (GreetRequest) returns (GreetReply) {
+        option (gapi.http) = {
+            post: "/greet"
+            use:  "auth"
+        };
    }
}
```

3. 生成pb.go

由于这里要引用到GAPI定义的proto,所以在生成的时候要将GAPI的上级目录加入到import path中。

```sh
protoc -I . -I 'path/to/gapi/root' --go_out=plugins=grpc:. *.proto
```