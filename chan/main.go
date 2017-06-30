package main

import (
	"context"
	"log"
	"time"
)

type Tester struct {
	ch  	chan int
	ctx 	context.Context
	cancel	context.CancelFunc
}

func NewTester() *Tester {
	var t Tester
	t.ch = make(chan int,1)
	t.ctx,t.cancel = context.WithCancel(context.Background())
	return &t
}

func main() {
	t:= NewTester()

	go ReadFromTester(t)

	ChangeChan(t)


	ChangeChanAndClosed(t)

	t.cancel()

	time.Sleep(100 * time.Millisecond)
}


func ReadFromTester(t *Tester){
	for {
		select {
		case <-t.ctx.Done():
			return
		case i,isOpen := <-t.ch:
			if !isOpen {
				log.Println("t.ch ： ",t.ch)
			}else {
				log.Println(" i : " ,i )
			}
		}
	}
}


func ChangeChan(t *Tester) {
	for i:=0; i < 10; i++ {
		t.ch <- i
	}

	// 可以工作，但是有数据丢失
	newChan := make(chan int, 1)
	t.ch = newChan

	for i:=10; i < 20; i++ {
		t.ch <- i
	}
}

func ChangeChanAndClosed(t *Tester) {
	for i:=0; i < 10; i++ {
		t.ch <- i
	}

	// 可以工作，但是有数据丢失
	newChan := make(chan int, 1)
	close(t.ch)
	t.ch = nil

	t.ch = newChan

	for i:=20; i < 30; i++ {
		t.ch <- i
	}
}