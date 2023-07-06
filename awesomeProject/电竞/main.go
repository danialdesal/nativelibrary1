package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

// fetchAndSaveJSON 从指定的 URL 获取 JSON 数据，将其解析后保存到指定的文件中
func fetchAndSaveJSON(url string, filename string) error {
	// 构造 HTTP 请求
	method := "GET"                               // 使用 GET 方法
	req, err := http.NewRequest(method, url, nil) // 创建 HTTP 请求对象
	if err != nil {
		return err
	}
	// 发送 HTTP 请求
	client := &http.Client{}    // 创建 HTTP 客户端对象
	resp, err := client.Do(req) // 发送 HTTP 请求
	if err != nil {
		return err
	}
	defer resp.Body.Close() // 关闭 HTTP 响应体

	// 读取 HTTP 响应
	body, err := ioutil.ReadAll(resp.Body) // 读取 HTTP 响应体的内容
	if err != nil {
		return err
	}

	// 解析 JSON
	var data interface{}                      // 定义一个空接口变量，用于保存解析后的 JSON 数据
	err = json.Unmarshal([]byte(body), &data) // 解析 JSON 数据
	if err != nil {
		return err
	}
	// 将 JSON 输出到文件
	file, err := os.Create(filename) // 创建指定文件名的文件
	if err != nil {
		return err
	}
	defer file.Close()                                 // 在函数返回时关闭文件
	jsonStr, err := json.MarshalIndent(data, "", "  ") // 将 JSON 数据格式化
	if err != nil {
		return err
	}
	file.Write(jsonStr)                       // 将格式化后的 JSON 数据写入文件
	fmt.Printf("JSON 数据已写入文件 %s\n", filename) // 输出提示信息
	return nil
}
func GenerateFiles() {
	// 获取接口数据并存储到文件
	err := fetchAndSaveJSON("https://lpl.qq.com/web201612/data/LOL_MATCH2_TEAM_TEAM422_INFO.js", "output1.json")
	if err != nil {
		fmt.Println(err) // 如果出错，则输出错误信息
		return           // 结束程序
	}

	err = fetchAndSaveJSON("https://lpl.qq.com/web201612/data/LOL_MATCH2_TEAM_TEAM57_INFO.js", "output2.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取更多接口数据并存储到文件
	err = fetchAndSaveJSON("https://lpl.qq.com/web201612/data/LOL_MATCH2_TEAM_TEAM1_INFO.js", "output3.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = fetchAndSaveJSON("https://lpl.qq.com/web201612/data/LOL_MATCH2_TEAM_TEAM7_INFO.js", "output4.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = fetchAndSaveJSON("https://lpl.qq.com/web201612/data/LOL_MATCH2_TEAM_TEAM2_INFO.js", "output5.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = fetchAndSaveJSON("https://lpl.qq.com/web201612/data/LOL_MATCH2_TEAM_TEAM29_INFO.js", "output6.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = fetchAndSaveJSON("https://lpl.qq.com/web201612/data/LOL_MATCH2_TEAM_TEAM4_INFO.js", "output7.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = fetchAndSaveJSON("https://lpl.qq.com/web201612/data/LOL_MATCH2_TEAM_TEAM9_INFO.js", "output8.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = fetchAndSaveJSON("https://lpl.qq.com/web201612/data/LOL_MATCH2_TEAM_TEAM587_INFO.js", "output9.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = fetchAndSaveJSON("https://lpl.qq.com/web201612/data/LOL_MATCH2_TEAM_TEAM6_INFO.js", "output10.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = fetchAndSaveJSON("https://lpl.qq.com/web201612/data/LOL_MATCH2_TEAM_TEAM11_INFO.js", "output11.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = fetchAndSaveJSON("https://lpl.qq.com/web201612/data/LOL_MATCH2_TEAM_TEAM8_INFO.js", "output12.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = fetchAndSaveJSON("https://lpl.qq.com/web201612/data/LOL_MATCH2_TEAM_TEAM42_INFO.js", "output13.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = fetchAndSaveJSON("https://lpl.qq.com/web201612/data/LOL_MATCH2_TEAM_TEAM438_INFO.js", "output14.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = fetchAndSaveJSON("https://lpl.qq.com/web201612/data/LOL_MATCH2_TEAM_TEAM685_INFO.js", "output15.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = fetchAndSaveJSON("https://lpl.qq.com/web201612/data/LOL_MATCH2_TEAM_TEAM41_INFO.js", "output16.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = fetchAndSaveJSON("https://lpl.qq.com/web201612/data/LOL_MATCH2_TEAM_TEAM12_INFO.js", "output17.json")
	if err != nil {
		fmt.Println(err)
		return
	}
}
func PutMongoDB() {
	// 创建 MongoDB 客户端
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查客户端连接状态
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to MongoDB!")
	}
	// 获取 MongoDB 集合
	collection := client.Database("LOL").Collection("lol")
	//读取 JSON 文件的内容，并将数据插入到 MongoDB 中
	for i := 1; i < 18; i++ {
		strI := strconv.Itoa(i)
		filename := "output" + strI + ".json"
		// 读取 JSON 文件的内容
		jsonBytes, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		// 解析 JSON 数据
		var data map[string]interface{}
		err = json.Unmarshal(jsonBytes, &data)
		if err != nil {
			log.Fatal(err)
		}
		// 将数据插入 MongoDB 中
		result, err := collection.InsertOne(context.Background(), data)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Inserted document with ID: %v\n", result.InsertedID)
		}
	}
	// 关闭 MongoDB 连接
	err = client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection to MongoDB closed.")
	}
}
func main() {
	GenerateFiles()
	PutMongoDB()
}
