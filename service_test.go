package transaction

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/suite"
)

type ServiceTest struct {
	suite.Suite
	txnService      Service
	datastoreMock   *MockDatastore
	transactionMock *Transaction
}

func (tests *ServiceTest) SetupTest() {
	datastoreMock := &MockDatastore{}
	tests.txnService = Service{
		datastore: datastoreMock,
	}
	tests.datastoreMock = datastoreMock
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	})
	tests.transactionMock = &Transaction{
		Amount:         100,
		SubmissionTime: time.Now(),
	}
}

func (tests *ServiceTest) TestCreateSuccess() {
	tests.datastoreMock.On("Create", *tests.transactionMock).Return(nil)
	actualTransaction, err := tests.txnService.Create(100)
	tests.Equal(tests.transactionMock, actualTransaction)
	tests.Nil(err)
}

func TestService(t *testing.T) {
	suite.Run(t, new(ServiceTest))
}
