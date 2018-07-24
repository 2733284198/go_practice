package main 

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"github.com/feixiao/log4go"
)

func main() {


	log4go.LoadConfiguration("./log.xml")

	gin.DefaultWriter = io.MultiWriter(os.Stdout,&GinLogger{})

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}