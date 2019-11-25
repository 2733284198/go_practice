package main


import (
	"github.com/feixiao/log4go"
)


type GinLogger struct {

}

// 显示行数不对
func (g *GinLogger) Write(p []byte) (n int, err error) {
	log4go.Info(string(p))
	return len(p), nil
}

