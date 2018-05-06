package database

import (
	"encoding/json"
	"errors"
)

type PaymentStatus int

const (
	Submitted PaymentStatus = 0
	Processed PaymentStatus = 1
	Confirmed PaymentStatus = 2
)

func (s PaymentStatus) String() string {
	switch s {
	case Submitted:
		return "submitted"
	case Processed:
		return "processed"
	case Confirmed:
		return "confirmed"
	default:
		return "unknown"
	}
}

func (p *PaymentStatus) Scan(value interface{}) error {
	paymentBytes, ok := value.([]byte)
	if !ok {
		return errors.New("Cannot convert to byte array")
	}
	switch string(paymentBytes) {
	case "submitted":
		*p = Submitted
	case "processed":
		*p = Processed
	case "confirmed":
		*p = Confirmed
	default:
		return errors.New("Cannot Scan the payment status")
	}
	return nil
}

func (s PaymentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
