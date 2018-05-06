package rpc

import "strconv"

func (r RPCService) GetBalance() (int, error) {
	var accountBalance string
	err := r.c.Call(&accountBalance, "eth_getBalance", r.account, "latest")
	if err != nil {
		return -1, err
	}
	b, err := strconv.ParseInt(accountBalance, 0, 64)
	return int(b), nil
}
