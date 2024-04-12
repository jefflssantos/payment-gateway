package entity

import (
	"errors"
	"regexp"
	"time"
)

type CreditCard struct {
	number          string
	name            string
	expirationMonth int
	expirationYear  int
	cvv             int
}

func NewCreditCard(number string, name string, expirationMonth int, expirationYear int, cvv int) *CreditCard {
	return &CreditCard{
		number:          number,
		name:            name,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		cvv:             cvv,
	}
}

func (creditCard *CreditCard) Validate() error {
	err := creditCard.ValidateNumber()
	if err != nil {
		return err
	}

	err = creditCard.ValidateExpirationDate()
	if err != nil {
		return err
	}

	return nil
}

func (creditCard *CreditCard) ValidateNumber() error {
	regex := regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`)

	if !regex.MatchString(creditCard.number) {
		return errors.New("invalid credit card number")
	}

	return nil
}

func (creditCard *CreditCard) ValidateExpirationDate() error {
	if creditCard.expirationYear < time.Now().Year() {
		return errors.New("invalid expiration year")
	}

	if creditCard.expirationMonth < int(time.Now().Month()) || creditCard.expirationMonth > 12 {
		return errors.New("invalid expiration month")
	}

	return nil
}
