package main

//
//import (
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"os"
//	"time"
//)
//
//func fetchAndSaveJSON(url string, filename string) error {
//	// 构造 HTTP 请求
//	method := "GET"
//	req, err := http.NewRequest(method, url, nil)
//	if err != nil {
//		return err
//	}
//
//	// 发送 HTTP 请求
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		return err
//	}
//	defer resp.Body.Close()
//
//	// 读取 HTTP 响应
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return err
//	}
//
//	// 解析 JSON
//	var data interface{}
//	err = json.Unmarshal([]byte(body), &data)
//	if err != nil {
//		return err
//	}
//
//	// 将 JSON 输出到文件
//	file, err := os.Create(filename)
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//
//	jsonStr, err := json.MarshalIndent(data, "", "  ")
//	if err != nil {
//		return err
//	}
//
//	file.Write(jsonStr)
//	fmt.Printf("JSON 数据已写入文件 %s\n", filename)
//	return nil
//}
//
//func main() {
//	// 设置定时器，每隔 1 分钟执行一次
//	ticker := time.NewTicker(1 * time.Minute)
//
//	for {
//		select {
//		case <-ticker.C:
//			// 获取接口数据并存储到文件
//			err := fetchAndSaveJSON("https://api.tjstats.com/auxiliary/public/basic?matchID=10331&BO=1&Match=LOL&Plat=douyu", "output1.json")
//			if err != nil {
//				fmt.Println(err)
//			}
//
//			err = fetchAndSaveJSON("https://api.tjstats.com/auxiliary/public/base?matchID=10341&BO=1&Match=LOL&Plat=douyu", "output2.json")
//			if err != nil {
//				fmt.Println(err)
//			}
//			// 获取更多接口数据并存储到文件
//			err = fetchAndSaveJSON("https://api.tjstats.com/auxiliary/public/overview?matchID=10341&BO=1&Match=LOL&Plat=douyu", "output3.json")
//			if err != nil {
//				fmt.Println(err)
//			}
//
//		}
//	}
//}
