package tcp

import (
	"net"
	"runtime"
	"strings"
	"log"
)


// 来自NSQ的TCP服务器实现

// TCP长连接的处理只需要实现Handle即可
type TCPHandler interface {
	Handle(net.Conn)
}

// 首先是监听端口，当有请求到来时开启一个goroutine去处理该链接请求
func TCPServer(listener net.Listener, handler TCPHandler) {
	log.Println("TCP: listening on %s", listener.Addr())

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
				log.Println("NOTICE: temporary Accept() failure - %s", nerr)
				runtime.Gosched()
				continue
			}
			// theres no direct way to detect this error because it is not exposed
			if !strings.Contains(err.Error(), "use of closed network connection") {
				log.Println("ERROR: listener.Accept() - %s", err)
			}
			break
		}
		go handler.Handle(clientConn)
	}

	log.Println("TCP: closing %s", listener.Addr())
}