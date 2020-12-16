package handlers

import (
	"context"
	"fmt"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api/apistruct"
	"net/http"
	"shitapi/log"
)

type apiT struct {
	Fullnode apistruct.FullNodeStruct
	Miner    apistruct.StorageMinerStruct
}

func Info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	authToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIl19.7w5TfZQZRgV4VzL2O_zYRmqUzy4uWIYHFLJ273kSLHs"
	headers := http.Header{"Authorization": []string{"Bearer " + authToken}}
	addr := "192.168.28.172:2345"

	var api apiT
	closer, err := jsonrpc.NewMergeClient(context.Background(), "ws://"+addr+"/rpc/v0", "Filecoin", []interface{}{&api.Fullnode.Internal, &api.Fullnode.CommonStruct.Internal, &api.Miner.Internal}, headers)
	if err != nil {
		log.Fatalf("connecting with lotus failed: %s", err)
	}
	defer closer()

	// Now you can call any API you're interested in.
	address, err := api.Miner.ActorAddress(context.Background())
	if err != nil {
		log.Fatalf("calling chain head: %s", err)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Miner:%s\n", address.String())
}
