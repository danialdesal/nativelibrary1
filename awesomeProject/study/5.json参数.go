package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func main() {
	byteData, _ := json.Marshal(map[string]any{
		"code": 0,
		"msg":  "成功",
		"data": map[string]string{
			"user_id": "1001",
		},
	})
	http.Post("http://127.0.0.1:7070/json", "application/json", bytes.NewBuffer(byteData))

	type Data struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	byteData, _ = json.Marshal(Data{Name: "淅栋", Age: 18})
	http.Post("http://127.0.0.1:7070/json", "application/json", bytes.NewBuffer(byteData))
}
