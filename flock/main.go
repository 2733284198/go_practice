package main

import (
	"os"
	"fmt"
	"sync"
	"time"
	. "github.com/feixiao/go_practice/flock/lockfile"
)

// golang的文件锁操作
// http://lihaoquan.me/2016/11/4/about-filelock.html
// windows & linux 上面的文件锁
// http://ikarishinjieva.github.io/blog/blog/2014/03/20/go-file-lock/

// 其他实现
// https://github.com/nightlyone/lockfile
// https://github.com/dickeyxxx/golock





func main() {
	test_file_path, _ := os.Getwd()
	locked_file := test_file_path

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			flock := NewFileLock(locked_file)
			err := flock.Lock()
			if err != nil {
				wg.Done()
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("output : %d\n", num)
			wg.Done()
		}(i)
	}
	wg.Wait()
	time.Sleep(2 * time.Second)
}