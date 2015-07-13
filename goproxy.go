package main

import (
	"github.com/winxxp/goproxy"
	"log"
	"net/http"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()

	log.Fatal(http.ListenAndServe(":80", proxy))
}
