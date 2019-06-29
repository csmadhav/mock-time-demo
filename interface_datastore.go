package transaction

// Datastore contract for db related operations
type Datastore interface {
	Create(txn Transaction) error
}
