package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	//serverHandler := &SimpleServerHandler{}
	//
	//rpcServer := jsonrpc.NewServer()
	//rpcServer.Register("SimpleServerHandler", serverHandler)
	r := mapRoutes()
	srv := &http.Server{
		Handler:           r,
		Addr:              "0.0.0.0:8000",
		WriteTimeout:      15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
