package main

import "net/http"

func main() {
	http.Get("http://127.0.0.1:7070/ping")
}
