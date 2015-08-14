package main

import (
	"github.com/winxxp/goproxy"
	"log"
	"net/http"
	"os"
)

func main() {
	httpHost := os.Getenv("HOST")
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "80"
	}

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			log.Println("--1--", r.URL.Path)
			return r, nil
		})
	log.Printf("proxy listening on %s:%s\n", httpHost, httpPort)
	log.Fatal(http.ListenAndServe(httpHost+":"+httpPort, proxy))
}
