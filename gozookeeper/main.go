package main

import (
	zk "github.com/samuel/go-zookeeper/zk"
	"time"
	"log"
)

var server []string = []string{"172.28.104.225:32770"}
var rootPath string = "/frank/test"
func main() {
	conn, _, err := zk.Connect(server, time.Second)
	if err != nil {
		log.Println(err)
		return
	}
	conn.Delete(rootPath,-1)
	path, err := conn.Create(rootPath, []byte("xxxx"), 0, zk.WorldACL(zk.PermAll))
	if err != nil {
		log.Println(err)
		return
	}
	defer  conn.Delete(rootPath,-1)

	log.Println("path : ",path)


	_ ,_,statChan,err := conn.GetW(rootPath)
	if err != nil {
		log.Println(err)
		return
	}

	conn.Set(rootPath,[]byte("11111"),-1)
	for {
		select {
		case s,isOpen := <-statChan:
			// 只通知一次
			if !isOpen {
				log.Println("statChan Closed")
				return
			}else {
				log.Printf("%+v",s)
			}
		}
	}

}