## 持续递交Golang项目

#### 安装Jenkins

#### 安装Go Plugin

+ 参考资料[《Go Plugin》](https://wiki.jenkins.io/display/JENKINS/Go+Plugin)


#### 配置项目信息

```shell
export GOPATH=$WORKSPACE
export PATH=$GOPATH/bin:$PATH

# 编译环境需要安装govendor和git
cd ${GOPATH}/src/hello
govendor sync
cd ${GOPATH}/src/hello 
chmod 755 build.sh
./build.sh  # hello目录下面的编译脚本

Delivery_Path="somewhere"
mv ${GOPATH}/src/hello/*tar.gz ${Delivery_Path}
```

#### 集成静态代码检测

+ [《Jenkins 之 Go 项目编译与代码静态检查》](https://www.testwo.com/blog/7920)

  ​

#### 参考资料

+ [《Build Golang Project in Jenkins》](https://smalltowntechblog.com/2014/11/30/build-golang-project-in-jenkins/)
+ [《Jenkins 之 Go 项目编译与代码静态检查》](https://www.testwo.com/blog/7920)
+ [《在 Jenkins 跑 Golang 測試》](https://blog.wu-boy.com/2016/08/golang-tesing-on-jenkins/)
+ [《Automate cross platform Golang builds with Jenkins》](https://medium.com/@reynn/automate-cross-platform-golang-builds-with-jenkins-ef7b07f1366e)
+ [《Golang X CI X CD》](https://kkc.github.io/2016/07/03/golang-ci/)
+ [《 Jenkins: Now with more Gopher》](http://www.asciiarmor.com/post/99010893761/jenkins-now-with-more-gopher)