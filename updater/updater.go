package updater

import (
	"time"

	"github.com/arjunrn/ethpay/database"
	"github.com/arjunrn/ethpay/rpc"
	log "github.com/sirupsen/logrus"
)

type Updater struct {
	d *database.DBService
	c *rpc.RPCService
}

func NewUpdater(db *database.DBService, c *rpc.RPCService) Updater {
	return Updater{d: db, c: c}
}

func (u Updater) Run() {
	sleepD, err := time.ParseDuration("1m")
	log.Debugf("Sleep Duration: %v", sleepD)
	if err != nil {
		panic(err)
	}
	for {
		err := u.updateTransactions()
		if err != nil {
			log.Error(err)
		}
		time.Sleep(sleepD)
	}
}

const confirmationBlocks = 5

func (u Updater) updateTransactions() error {
	lastestBlock, err := u.c.CurrentBlock()
	if err != nil {
		return err
	}
	log.Debugf("Latest Block %d", lastestBlock)
	pendingPayments, err := u.d.GetPendingPayments()
	log.Debugf("Number of pending payments: %d", len(pendingPayments))
	log.Info(pendingPayments)
	if err != nil {
		return err
	}
	for _, p := range pendingPayments {
		log.Infof("Checking status of payment %v", p)
		if p.PaymentStatus == database.Submitted {
			tx, err := u.c.GetTransactionByHash(p.TransactionHash)
			if err != nil {
				return err
			}
			if tx.IsPending() {
				log.Debugf("Transaction %s not mined yet. Continuing", p.TransactionHash)
				continue
			}
			bn, err := tx.BlockNumberInt()
			if err != nil {
				continue
			}
			log.Debugf("TX Block Number: %d ", bn)
			if bn != 0 {
				err = u.d.PaymentProcessed(p.TransactionHash, bn)
				if err != nil {
					return err
				}
			}
		} else if p.PaymentStatus == database.Processed {
			blockDepth := lastestBlock - p.BlockNumber
			if blockDepth >= confirmationBlocks {
				log.Infof("Transaction %s is %d blocks deep. Marking as confirmed", p.TransactionHash, blockDepth)
				err = u.d.MarkConfirmed(p)
				if err != nil {
					return err
				}
			} else {
				log.Debugf("Transaction %s still needs %d more confirmations. Skipping confirmation..", p.TransactionHash, confirmationBlocks-blockDepth)
			}
		}
	}
	return nil
}
