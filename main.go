package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arjunrn/ethpay/api"
	"github.com/arjunrn/ethpay/cmd"
	"github.com/arjunrn/ethpay/database"
	"github.com/arjunrn/ethpay/rpc"
	"github.com/arjunrn/ethpay/updater"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type response struct {
	JsonRpc string          `json:"jsonrpc"` // Version of the JSON RPC protocol, always set to 2.0
	Id      int             `json:"id"`      // Auto incrementing ID number for this request
	Error   json.RawMessage `json:"error"`   // Any error returned by the remote side
	Result  json.RawMessage `json:"result"`  // Whatever the remote side sends us in reply
}

func main() {
	log.SetLevel(log.DebugLevel)

	// account := "0x4673bda0a917e71db8d199ef92382d203d531c3b"
	// webAddr := "http://127.0.0.1:8545"
	// connString := "user=ethpay dbname=ethpay password=ethpay sslmode=disable"
	arguments := cmd.GetArguments()
	log.Debugf("Arguments: %v", arguments)
	rpcService, err := rpc.NewRPCService(arguments.GethAddr, arguments.EthAccount)

	if err != nil {
		panic(err)
	}
	defer rpcService.Close()
	blockCount, err := rpcService.CurrentBlock()
	if err != nil {
		panic(err)
	}
	log.Debugf("Block Count: %d", blockCount)

	balance, err := rpcService.GetBalance()
	if err != nil {
		panic(err)
	}
	log.Infof("Account Balance: %d", balance)

	dbService, err := database.NewDBService(arguments.PostgresConn)
	if err != nil {
		panic(err)
	}

	u := updater.NewUpdater(dbService, rpcService)
	go u.Run()
	router := api.NewRouter(rpcService, dbService, arguments.EthAccount)
	port := 8080
	log.Debugf("Starting API on port %d", port)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
