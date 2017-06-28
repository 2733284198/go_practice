package test

// Golang单元测试之httptest使用
// http://blog.csdn.net/lavorange/article/details/73369153?utm_source=itdadao&utm_medium=referral

import (
	"net/http"
	"fmt"
)

const (
	ADDRESS = "shanghai"
)

type Person struct {
	Name string `json:"name"`
	Address string `json:"address"`
	Age int `json:"age"`
}

func GetInfo(api string) ([]Person, error) {
	url := fmt.Sprintf("%s/person?addr=%s", api, ADDRESS)
	resp, err := http.Get(url)

	if err != nil {
		return []Person{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return []Person{}, fmt.Errorf("get info didn’t respond 200 OK: %s", resp.Status)
	}

	return nil, nil
}