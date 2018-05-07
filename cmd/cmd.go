package cmd

import (
	"fmt"
	"os"
)

type CommandArguments struct {
	GethAddr     string
	PostgresConn string
	EthAccount   string
}

func (a CommandArguments) String() string {
	return fmt.Sprintf("geth_addr=%s postgres_conn=\"%s\" ethereum_account=%s", a.GethAddr, a.PostgresConn, a.EthAccount)
}

func GetArguments() CommandArguments {
	if len(os.Args) != 4 {
		fmt.Printf("Usage: %s <geth_addr> <postgres_conn> <ethereum_account>", os.Args[0])
		os.Exit(1)
	}
	var arguments CommandArguments
	arguments.GethAddr = os.Args[1]
	arguments.PostgresConn = os.Args[2]
	arguments.EthAccount = os.Args[3]
	return arguments
}
