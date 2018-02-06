package main

// http://xorm.io/docs/

import (
	"math"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// golang xorm框架的使用
// http://www.cnblogs.com/guhao123/p/4159688.html
// xorm常用编程方法总结
// http://blog.csdn.net/wdy_yx/article/details/52687667

var engine *xorm.Engine

// xorm会根据成员变量的大小写生成数据库列名
// 	ServerID ——>  server_i_d
//  ServerId ——>  server_id
type User struct {
	Id   uint64
	Name string
}

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:root@tcp(172.28.104.222:3306)/test")
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
	user.Id = math.MaxUint64/2 + 10
	user.Name = "test"
	// 需要插入指定数据库表可以先通过Table()函数获取到数据库表
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

	user.Id = math.MaxUint64/2 + 10
	user.Name = "frank"

	_, err = engine.Where("Id=? and Name=?", user.Id, "fx").Update(&user)
	if err != nil {
		log.Println(err)
	}

	err = engine.Find(&pUsers)
	if err != nil {
		log.Println(err)
	}
	for k, v := range pUsers {
		log.Println(k, v)
	}
}
