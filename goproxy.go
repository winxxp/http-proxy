package main

import (
	"github.com/winxxp/goproxy"
	"os"
	"log"
	"net/http"
)

func main() {
	httpHost := os.Getenv("HOST")
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "80"
	}

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true

	log.Printf("proxy listening on %s:%s\n", httpHost, httpPort)
	log.Fatal(http.ListenAndServe(httpHost+":"+httpPort, proxy))
}
