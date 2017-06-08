package main

// go get github.com/pkg/errors
// Golang错误处理最佳方案	https://gocn.io/article/348

import (
	"log"
	"os"

	"github.com/pkg/errors"
)

func main() {
	_, err := os.Open("not_exsited_file.txt")
	if err != nil {
		log.Println(errors.Wrap(err, "read failed"))
		log.Println(errors.WithMessage(err, "open failed"))
		// Cause returns the underlying cause of the error, if possible.
		// An error value has a cause if it implements the following interface:
		log.Println(errors.Cause(err))

	}

	err1 := errors.New("whoops")
	err2 := errors.Wrap(err1, "111")
	log.Println(errors.WithStack(err2))
}
