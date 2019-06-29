package transaction

import (
	"fmt"
	"time"
)

// Service struct will define all dependencies for Transaction Service
type Service struct {
	datastore Datastore
}

// Create func creates a transaction
func (svc *Service) Create(amount int) (*Transaction, error) {
	transaction := Transaction{
		SubmissionTime: time.Now(),
		Amount:         amount,
	}
	err := svc.datastore.Create(transaction)
	if err != nil {
		fmt.Println(`error occured while creating in DB:`, err)
		return nil, err
	}
	return &transaction, nil
}
