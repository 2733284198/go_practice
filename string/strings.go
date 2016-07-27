package main

import (
	"strings"
	"unicode"
	"os"
	"log"
	"io/ioutil"
)

func main(){
	file, err := os.Open("./pg-moby_dick.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data, _:= ioutil.ReadAll(file)

	words := strings.FieldsFunc(string(data),func(ch rune) bool {
		return !unicode.IsLetter(ch) // 任何不是字母的字符都是分隔符
	})

	for _, word := range words {
		log.Println(word)
	}
}
