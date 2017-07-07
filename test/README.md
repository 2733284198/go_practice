## golang单元测试整理


#### golang基本的单元测试
+ go语言的单元测试采用内置的测试框架,通过引入testing包以及go test来提供测试功能。
+ _test.go为后缀名的源文件被go test认定为测试文件，这些文件不包含在go build的代码构建中,而是单独通过 go test来编译，执行。

#### 使用gotest生成表格驱动的测试用例
+ [gotests详细介绍](https://github.com/cweill/gotests)

+ gotests使用

  +  创建slice.go

  + 使用gotests根据slice.go的内容创建slice_test.go

    ```shell
    gotests -all slice.go -w slice_test.go
    ```

    生成代码类似如下内容：

    ```go
    func TestAdd(t *testing.T) {
    	type args struct {
    		s []string
    		a string
    	}
    	tests := []struct {
    		name string
    		args args
    		want []string
    	}{
    	// TODO: Add test cases.
    	}
    	for _, tt := range tests {
    		t.Run(tt.name, func(t *testing.T) {
    			if got := Add(tt.args.s, tt.args.a); !reflect.DeepEqual(got, tt.want) {
    				t.Errorf("Add() = %v, want %v", got, tt.want)
    			}
    		})
    	}
    }
    ```

#### mock的使用实践

+ [gomock](https://github.com/golang/mock)	
+ [gomock sample](https://github.com/golang/mock/tree/master/sample)
+ [testify mock](https://github.com/stretchr/testify#mock-package)
+ [mockery](https://github.com/vektra/mockery)  自动生成mock代码

#### HTTP服务单元测试

#### 代码覆盖率 

### 参考资料
+ [Testify - Thou Shalt Write Tests](https://github.com/stretchr/testify)
+ [Testing Your (HTTP) Handlers in Go](https://elithrar.github.io/article/testing-http-handlers-go/)
+ [Golang单元测试之httptest使用](http://blog.csdn.net/lavorange/article/details/73369153?utm_source=itdadao&utm_medium=referral)
+ [云平台下 Go 语言单元测试实践](http://www.golangtc.com/t/582be071b09ecc08ce0003b3)
+ [Basic testing patterns in Go](https://medium.com/agrea-technogies/basic-testing-patterns-in-go-d8501e360197)
+ [Mocking dependencies in Go](https://medium.com/agrea-technogies/mocking-dependencies-in-go-bb9739fef008)
