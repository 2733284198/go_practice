package main

// go get github.com/pkg/errors
// Golang错误处理最佳方案	https://gocn.io/article/348

import (
	"github.com/pkg/errors"
	"fmt"
	"log"
)

func main() {
	// 创建新的错误值
	err := errors.New("whoops")
	fmt.Println(err)         // 简单的输出错误信息
	fmt.Printf("%+v\n\n", err) // 错误信息加堆栈

	log.Println("***********************errors.WithMessage*************************************")
	err = f1()
	err1 := errors.WithMessage(err, "error in main") // 在原来的错误值之上添加额外的信息

	fmt.Printf("%+v\n\n", err) 	// 含有f1的调用信息
	fmt.Printf("%+v\n\n", err1) // err和err1的区别就是比err在最下方多了"error in main"


	log.Println("***********************errors.WithStack*************************************")
	err = f1()
	err1 = errors.WithStack(err) // 在原来的错误值之上添加额外的信息
	fmt.Printf("%+v\n\n", err)
	fmt.Printf("%+v\n\n", err1)

	log.Println("***********************errors.Cause*************************************")
	err = f1()
	err1 = errors.WithMessage(err, "error in main") // 在原来的错误值之上添加额外的信息
	fmt.Printf("%+s\n\n", errors.Cause(err1))
}


func f1() error {
	return errors.New("error in f1")
}
