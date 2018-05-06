package rpc

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	log "github.com/sirupsen/logrus"
)

type transactionParams struct {
	From  string
	To    string
	Value string
}

const ethInWei = float64(1000000000000000000)

func ethToWei(ether float64) *big.Int {
	if ether < 1 {
		return big.NewInt(int64(ether * ethInWei))
	} else {
		var result big.Int
		result.Mul(big.NewInt(int64(ether)), big.NewInt(int64(ethInWei)))
		return &result
	}
}

func (r RPCService) MakePayment(ether float64, recepient string) (string, error) {
	var transactionHash string
	weis := ethToWei(ether)
	log.Debugf("Weis: %v", weis)
	params := transactionParams{
		From:  r.account,
		To:    recepient,
		Value: hexutil.EncodeBig(weis),
	}
	if err := r.c.Call(&transactionHash, "eth_sendTransaction", params); err != nil {
		return "", nil
	}
	return transactionHash, nil
}
