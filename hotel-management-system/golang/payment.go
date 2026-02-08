package main

import "fmt"

type Payment interface {
	ProcessPayment(amount float64) error
}

type CashPayment struct{}

func NewCashPayment() *CashPayment {
	return &CashPayment{}
}

func (p *CashPayment) ProcessPayment(amount float64) error {
	// Process cash payment
	fmt.Println("Cash payment successful")
	return nil
}

type CreditCardPayment struct{}

func NewCreditCardPayment() *CreditCardPayment {
	return &CreditCardPayment{}
}

func (p *CreditCardPayment) ProcessPayment(amount float64) error {
	// Process credit card payment
	fmt.Println("CreditCard payment successful")
	return nil
}
