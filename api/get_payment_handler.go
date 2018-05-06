package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/arjunrn/ethpay/database"
	"github.com/gorilla/mux"
)

type GetPaymentHandler struct {
	d *database.DBService
}

func NewGetPaymentHandler(service *database.DBService) GetPaymentHandler {
	return GetPaymentHandler{d: service}
}

func (h GetPaymentHandler) ProcessRequest(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	var transactionHash string
	var ok bool
	if transactionHash, ok = vars["transaction_hash"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		return errors.New("Transaction Hash Not Present")
	}
	payment, err := h.d.GetPayment(transactionHash)
	if err == sql.ErrNoRows {
		return NotFoundError{model: "transaction"}
	}
	encoder := json.NewEncoder(w)
	return encoder.Encode(payment)
}
