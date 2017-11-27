package main

import (
	"container/list"
	"fmt"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/utils"
)

func main() {
	SampleList()

	list := arraylist.New()
	list.Add("a")                     // ["a"]
	list.Add("c", "b")                // ["a","c","b"]
	list.Sort(utils.StringComparator) // ["a","b","c"]
	fmt.Println(list)
}

// SampleList
func SampleList() {
	l := list.New() // 创建一个新的list
	// 添加5个元素
	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) // 输出list的值,01234
	}
	fmt.Println("")
	fmt.Println(l.Front().Value) // 输出首部元素的值,0
	fmt.Println(l.Back().Value)  // 输出尾部元素的值,4
	l.InsertAfter(6, l.Front())  // 首部元素之后插入一个值为10的元素
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) // 输出list的值,061234
		fmt.Print(" ")
	}
	fmt.Println(" ")
	l.MoveBefore(l.Front().Next(), l.Front()) // 首部两个元素位置互换
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) // 输出list的值,601234
		fmt.Print(" ")
	}
	fmt.Println("")
	l.MoveToFront(l.Back()) // 将尾部元素移动到首部
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) // 输出list的值,460123
		fmt.Print(" ")
	}
	fmt.Println("")
	l2 := list.New()
	l2.PushBackList(l) // 将l中元素放在l2的末尾
	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) // 输出l2的值,460123
		fmt.Print(" ")
	}
	fmt.Println(" ")
	fmt.Println(l.Len()) // 0
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) // 输出list的值,无内容
		fmt.Print(" ")
	}
}
