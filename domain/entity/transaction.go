package entity

import "errors"

const (
	REJECTED = "rejected"
	APPROVED = "approved"
)

type Transaction struct {
	Id           string
	AccountId    string
	Amount       float64
	CreditCard   CreditCard
	Status       string
	ErrorMessage string
}

func NewTransaction(id string, accountId string, amount float64, creditCard CreditCard) *Transaction {
	return &Transaction{
		Id:         id,
		AccountId:  accountId,
		Amount:     amount,
		CreditCard: creditCard,
	}
}

func (transaction *Transaction) Validate() error {
	if transaction.Amount <= 0 {
		return errors.New("invalid amount for this transaction")
	}

	if transaction.Amount >= 1000 {
		return errors.New("account do not have limit for this transaction")
	}

	return nil
}
