package handlers

import (
	"context"
	"fmt"
	"net/http"
	"shitapi/log"
	"shitapi/lotus"
)

func Info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	address, err := lotus.Client.Miner.ActorAddress(context.Background())
	if err != nil {
		log.Fatalf("calling chain head: %s", err)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Miner:%s\n", address.String())
}
