package rpc

import "strconv"

type Transaction struct {
	BlockNumber string `json:"blockNumber"`
}

func (t Transaction) BlockNumberInt() (int, error) {
	bn, err := strconv.ParseInt(t.BlockNumber, 0, 64)
	if err != nil {
		return -1, err
	}
	return int(bn), nil
}
func (t Transaction) IsPending() bool {
	return t.BlockNumber == ""
}

func (r RPCService) GetTransactionByHash(transactionHash string) (*Transaction, error) {
	var tx Transaction
	err := r.c.Call(&tx, "eth_getTransactionByHash", transactionHash)
	if err != nil {
		return nil, err
	}
	return &tx, nil
}
