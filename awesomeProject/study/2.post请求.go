package main

import (
	"bytes"
	"net/http"
	"net/url"
)

func main() {
	//application/x-www-form-urlencoded
	param := url.Values{}
	param.Add("name", "淅栋")
	http.PostForm("http://127.0.0.1:7070/form", param)

	http.Post(
		"http://127.0.0.1:7070/json",
		"application/json",
		bytes.NewBuffer([]byte(`{"code":0,"name":"淅栋"}`)),
	)

	//自己构建请求对象
	request, _ := http.NewRequest("GET", "http://127.0.0.1:7070/get", nil)
	http.DefaultClient.Do(request)
	request, _ = http.NewRequest("POST", "http://127.0.0.1:7070/post", nil)
	http.DefaultClient.Do(request)
	request, _ = http.NewRequest("PUT", "http://127.0.0.1:7070/put", nil)
	http.DefaultClient.Do(request)
	request, _ = http.NewRequest("DELETE", "http://127.0.0.1:7070/delete", nil)
	http.DefaultClient.Do(request)
}
