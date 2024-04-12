package process_transaction

import (
	"github.com/jefflssantos/imersao/gateway/domain/entity"
	"github.com/jefflssantos/imersao/gateway/domain/repository"
)

type ProcessTransaction struct {
	Repository  repository.TransactionRepository
	Transaction entity.Transaction
}

func NewProcessTransaction(repository repository.TransactionRepository, transaction entity.Transaction) *ProcessTransaction {
	return &ProcessTransaction{
		Repository:  repository,
		Transaction: transaction,
	}
}

func (processTransaction *ProcessTransaction) Process() TransactionDtoOutput {
	creditCardError := processTransaction.Transaction.CreditCard.Validate()
	if creditCardError != nil {
		return processTransaction.Response(entity.REJECTED, creditCardError.Error())
	}

	transactionError := processTransaction.Transaction.Validate()
	if transactionError != nil {
		return processTransaction.Response(entity.REJECTED, transactionError.Error())
	}

	return processTransaction.Response(entity.APPROVED, "")
}

func (processTransaction *ProcessTransaction) Response(status string, errorMessage string) TransactionDtoOutput {
	processTransaction.Repository.Insert(
		processTransaction.Transaction.Id,
		processTransaction.Transaction.AccountId,
		processTransaction.Transaction.Amount,
		status,
		errorMessage,
	)

	return TransactionDtoOutput{
		Id:           processTransaction.Transaction.Id,
		Status:       status,
		ErrorMessage: errorMessage,
	}
}
