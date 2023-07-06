package main

import "net/http"

func main() {
	request, _ := http.NewRequest("GET", "http://127.0.0.1:7070/head", nil)
	request.AddCookie(&http.Cookie{Name: "age", Value: "123"})
	request.AddCookie(&http.Cookie{Name: "name", Value: "123"})
	http.DefaultClient.Do(request)
}
