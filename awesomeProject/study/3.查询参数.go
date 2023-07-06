package main

import (
	"net/http"
	"net/url"
)

func main() {
	http.Get("http://127.0.0.1:7070/query?name=淅栋&age=12")

	param := url.Values{}
	param.Add("name", "淅栋")
	param.Add("age", "18")
	request, _ := http.NewRequest("GET", "http://127.0.0.1:7070/query", nil)
	request.URL.RawQuery = param.Encode()
	http.DefaultClient.Do(request)

	http.Get("http://127.0.0.1:7070/query?" + param.Encode())
}
