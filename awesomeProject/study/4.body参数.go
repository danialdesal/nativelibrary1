package main

import (
	"bytes"
	"net/http"
	"net/url"
)

func main() {
	param := url.Values{}
	param.Set("name", "淅栋")
	http.Post("http://127.0.0.1:7070/form", "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(param.Encode())))
	http.Post("http://127.0.0.1:7070/form", "multipart/form-data m", bytes.NewBuffer([]byte(param.Encode())))
}
