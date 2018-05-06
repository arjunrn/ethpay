package database

import (
	"fmt"
	"time"
)

type Payment struct {
	TransactionHash string        `json:"transactionHash"`
	LastUpdated     time.Time     `json:"lastUpdated"`
	PaymentStatus   PaymentStatus `json:"paymentStatus"`
	BlockNumber     int           `json:"blockNumber"`
}

func (p Payment) String() string {
	return fmt.Sprintf("%s - %s", p.TransactionHash, p.PaymentStatus)
}

func ParsePaymentStatus(paymentString string) (PaymentStatus, error) {
	switch paymentString {
	case "submitted":
		return Submitted, nil
	case "processed":
		return Processed, nil
	case "confirmed":
		return Confirmed, nil
	default:
		return -1, fmt.Errorf("Failed to parse payment with value: %s", paymentString)
	}
}

func (s DBService) NewPayment(transactionHash string) error {
	timestamp := time.Now()
	_, err := s.db.Exec("INSERT INTO payments (transaction_hash, last_updated, status) VALUES ($1, $2, $3)",
		transactionHash, timestamp, Submitted.String())
	return err
}

func (s DBService) GetPayment(transactionHash string) (*Payment, error) {
	row := s.db.QueryRow("SELECT last_updated, status, block_number FROM payments WHERE transaction_hash = $1", transactionHash)
	var p Payment
	var paymentStatus []byte
	p.TransactionHash = transactionHash
	err := row.Scan(&p.LastUpdated, &paymentStatus, &p.BlockNumber)
	if err != nil {
		return nil, err
	}
	status, err := ParsePaymentStatus(string(paymentStatus))
	if err != nil {
		return nil, err
	}
	p.PaymentStatus = status
	return &p, nil
}
