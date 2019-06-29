package transaction

import (
	"testing"
	"time"

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
	tests.transactionMock = &Transaction{
		Amount:         100,
		SubmissionTime: time.Now(),
	}
}

func (tests *ServiceTest) TestCreateSuccess() {
	tests.datastoreMock.On("Create", tests.transactionMock).Return(nil)
	actualTransaction, err := tests.txnService.Create(100)
	tests.Equal(tests.transactionMock, actualTransaction)
	tests.Nil(err)
}

func TestService(t *testing.T) {
	suite.Run(t, new(ServiceTest))
}
