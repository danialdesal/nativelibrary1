package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 构造 HTTP 请求
	url := "https://api.tjstats.com/auxiliary/public/overview?matchID=10341&BO=1&Match=LOL&Plat=douyu"
	method := "GET"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 发送 HTTP 请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// 读取 HTTP 响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 解析 JSON
	var data interface{}
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 输出 JSON
	jsonStr, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonStr))
}
