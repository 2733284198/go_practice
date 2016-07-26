package tcp

import (
	"net"
	"fmt"
	"log"
	"testing"
)

// go test -v ./tcpserver.go ./tcpserver_test.go

// 以EchoServer为例说明tcpserver的使用

// 定义类
type EchoServer struct {
	Listener net.Listener
}


// 创建新的EchoServer对象
func NewEchoServer(ip,port string)  *EchoServer {
	var server EchoServer
	var err error
	server.Listener,err = net.Listen("tcp",fmt.Sprintf("%s:%s",ip,port))
	if err != nil {
		panic(err)
	}
	return &server
}

// 实现Hande方法即可
func (s *EchoServer) Handle(conn net.Conn) {
	defer  conn.Close()
	data := make([]byte,1024*4)
	for{
		l, err := conn.Read(data)
		if err != nil  {
			log.Println(conn.RemoteAddr().String() + err.Error())
			break
		}
		conn.Write(data[:l])
	}
}

func TestTCPServer(t *testing.T) {
	server := NewEchoServer("0.0.0.0","8888")
	// 传入net.Listener和实现了Handle方法的类即完成了TCP服务器的开发
	TCPServer(server.Listener,server)
}