package main

import (
	"fmt"
	"log"
)

// http://studygolang.com/articles/2228

func main() {
	var ss []string

	//切片尾部追加元素append elemnt
	for i := 0; i < 10; i++ {
		ss = append(ss, fmt.Sprintf("s%d", i))
	}

	// 删除索引5的元素
	index := 5
	ss = append(ss[:index], ss[index+1:]...)
	log.Println(ss)

	for i := range ss {
		if i == 5 {
			ss = append(ss[:i], ss[i+1:]...)
		}
	}
	log.Println(ss)

	for i := range ss {
		if i == 5 {
			ss = append(ss[:i-1], ss[i:]...)
		}
	}
	log.Println(ss)
}
