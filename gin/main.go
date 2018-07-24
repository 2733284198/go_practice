package main 

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"github.com/feixiao/log4go"
	"github.com/gin-contrib/cors"
	"time"
)

func main() {


	log4go.LoadConfiguration("./log.xml")

	gin.DefaultWriter = io.MultiWriter(os.Stdout,&GinLogger{})

	r := gin.Default()

	config := cors.Config{
		AllowAllOrigins: true,
		AllowCredentials: false,
		MaxAge: 12 * time.Hour,
	}

	// "Origin,No-Cache, X-Requested-With, If-Modified-Since, Pragma, Last-Modified, Cache-Control, Expires, Content-Type, X-E4M-With, userId, token"
	config.AddAllowHeaders("token","x-access-token","x-url-path")

	r.Use(cors.New(config))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}