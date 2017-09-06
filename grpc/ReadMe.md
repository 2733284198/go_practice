## gRPC快速入门

#### 预备知识

+ gRPC需要Go 1.6版本以上

  ```shell
  $ go version
  go version go1.8 linux/amd64   # 使用的ubuntu上面的go1.8版本
  ```

#### 安装gRPC

  + 从github获取最新的开发版本

    ```shell
    go get google.golang.org/grpc  # 需要翻墙
    ```

  + 从github指定版本

    ```shell
    # 进入GOPATH目录，创建grpc目录
    cd ~/gohome/src/go_practice/grpc
    # 初始化
    govendor init 
    # 或者1.6.0版本
    govendor fetch google.golang.org/grpc@v1.6.0 
    ```

+ 安装protobuf

  ```shell
  go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

  # 为了安装命令行工具protoc
  https://github.com/google/protobuf/releases/tag/v3.4.0
  ```


#### 例子分析

+  进入example目录

  ```shell
  cd ~/gohome/src/google.golang.org/grpc/examples/helloworld
  ```

+ 项目结构

  ```shell
  .
  ├── greeter_client   	# 客户端
  │   └── main.go
  ├── greeter_server   	# 服务端
  │   └── main.go
  ├── helloworld 		 	# 协议定义
  │   ├── helloworld.pb.go  # 由helloworld.proto生成
  │   └── helloworld.proto  # 具体的协议定义
  └── mock_helloworld		# 测试代码
      ├── hw_mock.go
      └── hw_mock_test.go
  ```

  ​

#### 参考资料

+ [Go Quick Start](https://grpc.io/docs/quickstart/go.html)
+ [protocol-buffers](https://developers.google.com/protocol-buffers/)
+ [gRPC初接触](https://samael65535.github.io/2017-05-18/grpc_newb/)