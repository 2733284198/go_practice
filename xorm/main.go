package main

// http://xorm.io/docs/

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

type User struct {
	Id   int64
	Name string
}

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:12345678@tcp(172.28.104.225:32771)/test")
	if err != nil {
		log.Println(err)
		return
	}

	err = engine.Ping()
	if err != nil {
		log.Println(err)
		return
	}
	// 设置日志
	f, err := os.Create("sql.log")
	if err != nil {
		println(err.Error())
		return
	}
	engine.SetLogger(xorm.NewSimpleLogger(f))

	// 创建数据库表
	// 操作过程中没有表的概念，只有对象的概念，操作对象即操作相应的数据库表
	engine.CreateTables(User{})

	// 插入数据
	var user User
	user.Id = 1
	user.Name = "test"
	engine.Insert(&user)

	user.Id = 2
	user.Name = "fei"
	engine.Insert(&user)

	pUsers := make(map[int64]*User)
	err = engine.Find(&pUsers)
	if err != nil {
		log.Println(err)
	}
	for k, v := range pUsers {
		log.Println(k, v)
	}
}
