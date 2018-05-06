package api

import (
	"encoding/json"
	"net/http"

	"github.com/arjunrn/ethpay/database"
	"github.com/arjunrn/ethpay/rpc"
	log "github.com/sirupsen/logrus"
)

type CreatePaymentHandler struct {
	dbService  *database.DBService
	account    string
	rpcService *rpc.RPCService
}

type newPaymentResponse struct {
	TransactionHash string `json:"transactionHash"`
}

type PaymentInput struct {
	To    string  `json:"to"`
	Ether float64 `json:"ether"`
}

func NewCreatePaymentHandler(account string, r *rpc.RPCService, d *database.DBService) *CreatePaymentHandler {
	return &CreatePaymentHandler{account: account, rpcService: r, dbService: d}
}

func (h CreatePaymentHandler) ProcessRequest(w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var input PaymentInput
	if err := decoder.Decode(&input); err != nil {
		return err
	}
	transactionHash, err := h.rpcService.MakePayment(input.Ether, input.To)
	if err != nil {
		return err
	}
	log.Debugf("Transaction Hash: %s", transactionHash)
	err = h.dbService.NewPayment(transactionHash)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	return encoder.Encode(&newPaymentResponse{TransactionHash: transactionHash})
}
