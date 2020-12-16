package main

import (
	"log"
	"net/http"
	"shitapi/lotus"
	"time"
)

func main() {

	closer := lotus.Init()
	defer closer()

	r := mapRoutes()
	srv := &http.Server{
		Handler:           r,
		Addr:              "0.0.0.0:8000",
		WriteTimeout:      15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
