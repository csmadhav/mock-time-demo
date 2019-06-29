package transactiondb

import (
	"fmt"

	"github.com/csmadhav/mock-time-demo/transaction"
)

// Datastore struct for defining all dependencies for datastore
type Datastore struct {
	// something like db gorm.DB
}

// Create for creating a transaction
func (db *Datastore) Create(txn transaction.Transaction) error {
	fmt.Println(`Creating txn with data:`, txn)
	// your DB operations to transactions ...
	return nil
}
