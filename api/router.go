package api

import (
	"github.com/arjunrn/ethpay/database"
	"github.com/arjunrn/ethpay/rpc"

	"github.com/gorilla/mux"
)

func NewRouter(c *rpc.RPCService, s *database.DBService, account string) *mux.Router {
	r := mux.NewRouter()
	wPaymentHandler := NewWrappedErrorHandler(NewCreatePaymentHandler(account, c, s))
	wGetPaymentHandler := NewWrappedErrorHandler(NewGetPaymentHandler(s))
	wBalancedHandler := NewWrappedErrorHandler(NewBalanceHandler(account, c))
	r.Handle("/payment", wPaymentHandler).Methods("POST")
	r.Handle("/payment/{transaction_hash}", wGetPaymentHandler).Methods("GET")
	r.Handle("/balance", wBalancedHandler).Methods("GET")
	return r
}
