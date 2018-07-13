package main

import (
	"fmt"
	"encoding/json"
)


// Json Marshal Interface

type Animal interface {
    Speak() string
}

type Cat struct{
	Age int 
}
func (c Cat) Speak() string {
    return "cat"
}

type Dog struct{
	Age int 
}

func (d Dog) Speak() string {
    return "dog"
}


func main() {
	m:= make(map[int]Animal, 16)
	m[1] = &Dog{
		Age: 1,
	}
	buf, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Println(string(buf))
}