package main

import (
	"encoding/json"
	"os"
	"fmt"
)

func main() {
	// encoder的结果输出到标准输出
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)


	// encoder的结果输出到文件，注意文件的打开模式
	// https://segmentfault.com/a/1190000000376807
	file, err := os.OpenFile("test.json",os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fenc := json.NewEncoder(file)
	fenc.Encode(d)
}
