package main 

import (
	"time"
	"fmt"
	"context"
	sync "github.com/funny/debug/sync"
	//"fmt"
)

func main() {
		var lockA sync.RWMutex
		var lockB sync.RWMutex 
		ctx, _ := context.WithCancel(context.Background())

		go func(la,lb *sync.RWMutex,ctx context.Context) {
				la.Lock()
				fmt.Println("la.Lock()")
				lb.Lock()
				fmt.Println("lb.Lock()")
				<- ctx.Done()
		}(&lockA, &lockB, ctx)


		time.Sleep(3*time.Second)
		lockB.Lock()
		fmt.Println("lockB.Lock()")
		lockA.Lock()
		fmt.Println("lockA.Lock()")
		<- ctx.Done()
		//cancel()
}