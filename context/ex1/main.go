package main

import (
	"context"
	"fmt"
	"time"
)

// Go Context
// http://www.flysnow.org/2017/05/12/go-in-action-go-context.html?utm_source=tuicool&utm_medium=referral

/*
	context.Background() 返回一个空的Context，这个空的Context一般用于整个Context树的根节点。
	然后我们使用context.WithCancel(parent)函数，创建一个可取消的子Context，然后当作参数传给goroutine使用，这样就可以使用这个子Context跟踪这个goroutine。

	在goroutine中，使用select调用<-ctx.Done()判断是否要结束，如果接受到值的话，就可以返回结束goroutine了；如果接收不到，就会继续进行监控。
	那么是如何发送结束指令的呢？这就是示例中的cancel函数啦，它是我们调用context.WithCancel(parent)函数生成子Context的时候返回的，第二个返回值就是这个取消函数，它是CancelFunc类型的。
	我们调用它就可以发出取消指令，然后我们的监控goroutine就会收到信号，就会返回结束。

*/

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context){
		for {
			select {
			case <- ctx.Done():
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("运行中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	// 为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)

	// 注：只能通知退出，但是不能保证在main之前退出
	// defer fmt.Println("主goroutine退出中...")
}