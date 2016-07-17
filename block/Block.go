package main

import (
	"sync"
	"log"
	"time"
	"net/http"
	_ "net/http/pprof"
)

func main() {
        // http://localhost:8080/debug/pprof/
	go http.ListenAndServe(":8080", nil) // for pprof

	b := NewBlock()

	b.Start()
	go b.sendData(16)
	time.Sleep(1 * time.Second)

	b.Stop()
}


type Block struct {
	exitChan chan bool
	dataChan chan int
	wg  	 WaitGroupWrapper
}

func NewBlock() *Block {
	var b Block
	b.dataChan = make(chan int,16)
	b.exitChan = make(chan bool,1)
	return &b
}

func (b *Block) Start() {
	b.wg.Wrap(func(){b.handleData()})
}

func (b *Block) Stop() {
	log.Println("Block Stop")
	if b.exitChan != nil {
		close(b.exitChan)
		b.exitChan = nil
	}

	b.wg.Wait()

	if b.dataChan != nil {
		close(b.dataChan)
		b.dataChan = nil
	}


}

func (b *Block) sendData(max int ) {
	for i:=0; i<max; i++ {
		b.dataChan <- i
	}
}

func (b *Block) handleData() {
	for  {
		select {
		case i:= <-b.dataChan:
			log.Println(i)
			time.Sleep(200 * time.Millisecond)
		case <- b.exitChan:
			goto exit
		}
	}
	exit:
	log.Println("handleData Out")
}


type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		cb()
		w.Done()
	}()
}
