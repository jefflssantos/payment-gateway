package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTransaction(t *testing.T) {
	creditcard := NewCreditCard("5204996669957419", "Jefferson Santos", 12, time.Now().Year(), 123)

	transaction := NewTransaction("1", "1", 200, *creditcard)
	assert.Nil(t, transaction.Validate())

	transaction = NewTransaction("1", "1", 1000, *creditcard)
	assert.Equal(t, "account do not have limit for this transaction", transaction.Validate().Error())

	transaction = NewTransaction("1", "1", 0, *creditcard)
	assert.Equal(t, "invalid amount for this transaction", transaction.Validate().Error())

	transaction = NewTransaction("1", "1", -5, *creditcard)
	assert.Equal(t, "invalid amount for this transaction", transaction.Validate().Error())
}
