package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// 验证 syscall.SIGKILL 不能被捕获
// kill -9 pid
// kill -15 pid
func main() {
	ss := make(chan os.Signal)
	signal.Notify(ss, syscall.SIGKILL, syscall.SIGTERM) // 并不能捕获该信号

	log.Println("Wait...")
	<-ss
	log.Println("Done...")
}
