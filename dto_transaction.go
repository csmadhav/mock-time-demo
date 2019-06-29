package transaction

import "time"

// Transaction used to communicate between services and networks
type Transaction struct {
	ID             int
	Amount         int
	SubmissionTime time.Time
}
