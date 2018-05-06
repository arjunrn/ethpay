package database

func (d DBService) PaymentProcessed(transactionHash string, blockNumber int) error {
	_, err := d.db.Exec("UPDATE payments SET block_number = $1, status = $2 WHERE transaction_hash = $3", blockNumber, Processed.String(), transactionHash)
	return err
}

func (d DBService) MarkConfirmed(p Payment) error {
	_, err := d.db.Exec("UPDATE payments SET status = $1 WHERE transaction_hash = $2", Confirmed.String(), p.TransactionHash)
	return err
}
