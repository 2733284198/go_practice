## pkg/errors快速入门

[errors源码地址](https://github.com/pkg/errors)

#### 接口介绍

```go
func Cause(err error) error
func Errorf(format string, args ...interface{}) error
func New(message string) error
func WithMessage(err error, message string) error
func WithStack(err error) error
func Wrap(err error, message string) error
func Wrapf(err error, format string, args ...interface{}) error
```

#### New

创建新的错误值。

```go
err := errors.New("whoops")
fmt.Println(err)               // 简单的输出错误信息
fmt.Printf("%+v\n",err)        // 错误信息加堆栈

output：
whoops
whoops
main.main	/home/frank/gohome/src/github.com/feixiao/go_practice/errors/main.go:14
runtime.main
	/opt/go1_8/src/runtime/proc.go:185
runtime.goexit
	/opt/go1_8/src/runtime/asm_amd64.s:2197

```

#### WithMessage

在原来的错误上面添加新的错误内容(同时堆栈是函数的调用关系)。

```go
	err = f1()
	err1 := errors.WithMessage(err, "error in main") // 在原来的错误值之上添加额外的信息

	fmt.Printf("%+v\n\n", err) 	// 含有f1的调用信息
	fmt.Printf("%+v\n\n", err1) // err和err1的区别就是比err在最下方多了"error in main"
```

#### WithStack

在原来的错误值上面添加新的堆栈，新添加的堆栈信息跟之前的并列关系。

```go
	err = f1()
	err1 = errors.WithStack(err) // 在原来的错误值之上调用处的堆栈信息
	fmt.Printf("%+v\n\n", err)
	fmt.Printf("%+v\n\n", err1)
output：
error in f1
main.f1	/home/frank/gohome/src/github.com/feixiao/go_practice/errors/main.go:35
main.main	/home/frank/gohome/src/github.com/feixiao/go_practice/errors/main.go:27
runtime.main
	/opt/go1_8/src/runtime/proc.go:185
runtime.goexit
	/opt/go1_8/src/runtime/asm_amd64.s:2197

error in f1
main.f1	/home/frank/gohome/src/github.com/feixiao/go_practice/errors/main.go:35
main.main	/home/frank/gohome/src/github.com/feixiao/go_practice/errors/main.go:27
runtime.main
	/opt/go1_8/src/runtime/proc.go:185
runtime.goexit
	/opt/go1_8/src/runtime/asm_amd64.s:2197
main.main	/home/frank/gohome/src/github.com/feixiao/go_practice/errors/main.go:28
runtime.main
	/opt/go1_8/src/runtime/proc.go:185
runtime.goexit
	/opt/go1_8/src/runtime/asm_amd64.s:2197
```

#### Wrap/Wrapf

在原来的错误值上面添加信息和堆栈消息。

#### Cause

```
	err = f1()
	err1 = errors.WithMessage(err, "error in main") // 在原来的错误值之上添加额外的信息
	fmt.Printf("%+s\n\n", errors.Cause(err1))
output：
error in f1

```