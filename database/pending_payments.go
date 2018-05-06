package database

func (d DBService) GetPendingPayments() ([]Payment, error) {
	rows, err := d.db.Query("SELECT transaction_hash, status, last_updated, block_number FROM payments WHERE status = $1 OR status =$2", Submitted.String(), Processed.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	payments := make([]Payment, 0)
	for rows.Next() {
		var payment Payment
		err := rows.Scan(&payment.TransactionHash, &payment.PaymentStatus, &payment.LastUpdated, &payment.BlockNumber)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}
	return payments, nil
}
