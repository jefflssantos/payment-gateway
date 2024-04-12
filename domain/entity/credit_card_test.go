package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {
	creditCard := NewCreditCard("5204996669957419", "Jefferson Santos", 12, time.Now().Year(), 123)
	assert.Nil(t, creditCard.Validate())

	creditCard = NewCreditCard("0909009090909909", "Jefferson Santos", 12, time.Now().Year(), 123)
	assert.Equal(t, "invalid credit card number", creditCard.Validate().Error())
}

func TestCreditCardExpirationDate(t *testing.T) {
	creditCard := NewCreditCard("5204996669957419", "Jefferson Santos", 12, time.Now().Year(), 123)
	assert.Nil(t, creditCard.Validate())

	creditCard = NewCreditCard("5204996669957419", "Jefferson Santos", 13, time.Now().Year(), 123)
	assert.Equal(t, "invalid expiration month", creditCard.Validate().Error())

	creditCard = NewCreditCard("5204996669957419", "Jefferson Santos", 0, time.Now().Year(), 123)
	assert.Equal(t, "invalid expiration month", creditCard.Validate().Error())

	creditCard = NewCreditCard("5204996669957419", "Jefferson Santos", -1, time.Now().Year(), 123)
	assert.Equal(t, "invalid expiration month", creditCard.Validate().Error())

	creditCard = NewCreditCard("5204996669957419", "Jefferson Santos", time.Now().AddDate(0, -1, 0).Year(), time.Now().Year(), 123)
	assert.Equal(t, "invalid expiration month", creditCard.Validate().Error())

	creditCard = NewCreditCard("5204996669957419", "Jefferson Santos", 12, time.Now().AddDate(-1, 0, 0).Year(), 123)
	assert.Equal(t, "invalid expiration year", creditCard.Validate().Error())

	creditCard = NewCreditCard("5204996669957419", "Jefferson Santos", 12, 123, 123)
	assert.Equal(t, "invalid expiration year", creditCard.Validate().Error())
}
