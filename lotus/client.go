package lotus

import (
	"context"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api/apistruct"
	"net/http"
	"shitapi/log"
)

type apiT struct {
	Fullnode apistruct.FullNodeStruct
	Miner    apistruct.StorageMinerStruct
}

var Client apiT

func Init() jsonrpc.ClientCloser {
	authToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIl19.7w5TfZQZRgV4VzL2O_zYRmqUzy4uWIYHFLJ273kSLHs"
	headers := http.Header{"Authorization": []string{"Bearer " + authToken}}
	addr := "192.168.28.172:2345"

	closer, err := jsonrpc.NewMergeClient(context.Background(), "ws://"+addr+"/rpc/v0", "Filecoin", []interface{}{&Client.Fullnode.Internal, &Client.Fullnode.CommonStruct.Internal, &Client.Miner.Internal}, headers)
	if err != nil {
		log.Fatalf("connecting with lotus failed: %s", err)
	}

	return closer
}
