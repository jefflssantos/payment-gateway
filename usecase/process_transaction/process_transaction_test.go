package process_transaction

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jefflssantos/imersao/gateway/domain/entity"
	mock_repository "github.com/jefflssantos/imersao/gateway/domain/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestProcessTransaction_ValidTransaction(t *testing.T) {
	transactionInput := TransactionDtoInput{
		Id:                        "1",
		AccountId:                 "1",
		CreditCardNumber:          "5204996669957419",
		CreditCardName:            "Jefferson Santos",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    200,
	}

	creditCard := entity.NewCreditCard(
		transactionInput.CreditCardNumber,
		transactionInput.CreditCardName,
		transactionInput.CreditCardExpirationMonth,
		transactionInput.CreditCardExpirationYear,
		transactionInput.CreditCardCVV,
	)

	transaction := entity.NewTransaction(
		transactionInput.Id,
		transactionInput.AccountId,
		transactionInput.Amount,
		*creditCard,
	)

	expectedOutput := TransactionDtoOutput{
		Id:           "1",
		Status:       entity.APPROVED,
		ErrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(transaction.Id, transaction.AccountId, transaction.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	processTransactionResponse := NewProcessTransaction(repositoryMock, *transaction).Process()

	assert.Equal(t, expectedOutput, processTransactionResponse)
}

func TestProcessTransaction_InvalidCreditCard(t *testing.T) {
	transactionInput := TransactionDtoInput{
		Id:                        "1",
		AccountId:                 "1",
		CreditCardNumber:          "0000000000000000",
		CreditCardName:            "Jefferson Santos",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    200,
	}

	creditCard := entity.NewCreditCard(
		transactionInput.CreditCardNumber,
		transactionInput.CreditCardName,
		transactionInput.CreditCardExpirationMonth,
		transactionInput.CreditCardExpirationYear,
		transactionInput.CreditCardCVV,
	)

	transaction := entity.NewTransaction(
		transactionInput.Id,
		transactionInput.AccountId,
		transactionInput.Amount,
		*creditCard,
	)

	expectedOutput := TransactionDtoOutput{
		Id:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: creditCard.Validate().Error(),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(transaction.Id, transaction.AccountId, transaction.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	processTransactionResponse := NewProcessTransaction(repositoryMock, *transaction).Process()

	assert.Equal(t, expectedOutput, processTransactionResponse)
}

func TestProcessTransaction_InvalidTransactionAmount(t *testing.T) {
	transactionInput := TransactionDtoInput{
		Id:                        "1",
		AccountId:                 "1",
		CreditCardNumber:          "5204996669957419",
		CreditCardName:            "Jefferson Santos",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    1000,
	}

	creditCard := entity.NewCreditCard(
		transactionInput.CreditCardNumber,
		transactionInput.CreditCardName,
		transactionInput.CreditCardExpirationMonth,
		transactionInput.CreditCardExpirationYear,
		transactionInput.CreditCardCVV,
	)

	transaction := entity.NewTransaction(
		transactionInput.Id,
		transactionInput.AccountId,
		transactionInput.Amount,
		*creditCard,
	)

	expectedOutput := TransactionDtoOutput{
		Id:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "account do not have limit for this transaction",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(transaction.Id, transaction.AccountId, transaction.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	processTransactionResponse := NewProcessTransaction(repositoryMock, *transaction).Process()

	assert.Equal(t, expectedOutput, processTransactionResponse)
}
