## kratos快速入门

### 安装kratos

```shell
go get -u github.com/bilibili/kratos/tool/kratos
```

### 创建demo

```
cd $GOPATH/src
kratos new kratos-demo
cd kratos-demo/cmd 
go build 
./cmd -conf ../configs


# 访问
http://localhost:8000/kratos-demo/start
```

