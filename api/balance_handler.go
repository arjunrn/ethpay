package api

import (
	"encoding/json"
	"net/http"

	"github.com/arjunrn/ethpay/rpc"
	log "github.com/sirupsen/logrus"
)

type BalanceHandler struct {
	account string
	r       *rpc.RPCService
}

func NewBalanceHandler(account string, client *rpc.RPCService) BalanceHandler {
	return BalanceHandler{account: account, r: client}
}

type balanceResponse struct {
	Balance int    `json:"balance"`
	Account string `json:"account"`
}

func (h BalanceHandler) ProcessRequest(w http.ResponseWriter, r *http.Request) error {
	balance, err := h.r.GetBalance()
	if err != nil {
		return err
	}
	log.Debugf("Account Balance: %v", balance)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(balanceResponse{Account: h.account, Balance: int(balance)})
	return err
}
