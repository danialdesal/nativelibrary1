package main

import "net/http"

func main() {
	request, _ := http.NewRequest("GET", "http://127.0.0.1:7070/head", nil)
	request.Header.Set("age", "18")
	request.Header.Add("age", "3456")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.58")
	http.DefaultClient.Do(request)
}
