package main

import (
	"encoding/json"
	"os"
	"log"
	"fmt"
)
type KeyValue struct {
	Key   string
	Value string
}

func main() {
	// encoder的结果输出到标准输出
	enc := json.NewEncoder(os.Stdout)

	var v1 KeyValue
	v1.Key = "Hello"
	v1.Value = "2"

	var v2 KeyValue
	v2.Key = "World"
	v2.Value= "3"

	enc.Encode(v1)
	enc.Encode(v2)


	file, err := os.OpenFile("test.json",os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fenc := json.NewEncoder(file)
	fenc.Encode(v1)
	fenc.Encode(v2)

	/*
	{"Key":"Hello","Value":"2"}
	{"Key":"World","Value":"3"}*/

	// decoder
	var rfile *os.File
	rfile,err = os.Open("test.json")
	if err != nil {
		log.Fatal(err)
	}
	defer rfile.Close()

	fdec := json.NewDecoder(rfile)
	var kvs map[string]string = make(map[string]string,128)
	for {
		var kv KeyValue
		err = fdec.Decode(&kv)
		if err != nil {
			break
		}
		kvs[kv.Key] = kv.Value
	}
	fmt.Println(kvs)
	// map[World:3 Hello:2]


}
