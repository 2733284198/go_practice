package main 

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"github.com/feixiao/log4go"
	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/cors"
	"time"
)

// http://www.jyguagua.com/?p=3038
// https://segmentfault.com/q/1010000015287655

func main() {


	log4go.LoadConfiguration("./log.xml")

	gin.DefaultWriter = io.MultiWriter(os.Stdout,&GinLogger{})

	r := gin.Default()


	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type","token","x-access-token","x-url-path"},
		AllowAllOrigins: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	pprof.Register(r)

	r.OPTIONS("/ping",)
	r.Run(":8000") // listen and serve on 0.0.0.0:8000
}