package main

import (
	"encoding/json"
	"os"
	"log"
)

func main() {
	// encoder的结果输出到标准输出
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)


	file, err := os.OpenFile("test.json",os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fenc := json.NewEncoder(file)
	fenc.Encode(d)


	// decoder
	var rfile *os.File
	rfile,err = os.Open("test.json")
	if err != nil {
		log.Fatal(err)
	}
	defer rfile.Close()

	fdec := json.NewDecoder(rfile)
	var data map[string]int
	if err := fdec.Decode(&data); err != nil{
		log.Println(err)
	}else {
		log.Println(data)
	}


}
