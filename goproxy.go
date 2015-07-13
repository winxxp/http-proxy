package main

import (
	"github.com/winxxp/goproxy"
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
	
	log.Printf("proxy listening on %s:%s\n", httpHost, httpPort)
	log.Fatal(http.ListenAndServe(httpHost+":"+httpPort, proxy))
}
