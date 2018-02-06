package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"
	// app.Command 命令行参数
	//  执行的方法
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello friend!")
		return nil
	}

	app.Run(os.Args)
}
