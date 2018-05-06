package rpc

import (
	"strconv"

	"github.com/ethereum/go-ethereum/rpc"
	log "github.com/sirupsen/logrus"
)

type RPCService struct {
	account string
	c       *rpc.Client
}

func NewRPCService(connectionString, account string) (*RPCService, error) {
	client, err := rpc.Dial(connectionString)
	if err != nil {
		return nil, err
	}
	return &RPCService{c: client, account: account}, nil
}

func (r RPCService) Ping() error {
	var syncingResponse interface{}
	err := r.c.Call(&syncingResponse, "eth_syncing", nil)
	if err != nil {
		return err
	}
	log.Debugf("eth_syncing %v", syncingResponse)
	return nil
}

func (r RPCService) CurrentBlock() (int, error) {
	var blockCount string
	err := r.c.Call(&blockCount, "eth_blockNumber", nil)
	if err != nil {
		return -1, err
	}
	b, err := strconv.ParseInt(blockCount, 0, 64)
	if err != nil {
		return -1, err
	}
	return int(b), nil
}

func (r RPCService) Close() {
	r.c.Close()
}
