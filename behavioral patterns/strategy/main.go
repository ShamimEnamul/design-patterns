package main

import "fmt"

// Strategy

type PaymentStrategy interface {
	pay(amount int64)
}

// concrete strategy paypal

type PaypalPayment struct{}

func (p *PaypalPayment) pay(amount int64) {
	fmt.Printf("paypal payment done")
}

// concrete strategy creditcard

type CreditCardPayment struct{}

func (c *CreditCardPayment) pay(amount int64) {
	fmt.Printf("credit card payment doe")
}

// context

type PaymentContext struct {
	paymentStrategy PaymentStrategy
}

func (c *PaymentContext) pay(amount int64) {
	c.paymentStrategy.pay(amount)
}
func (c *PaymentContext) setPaymentStrategy(strategy PaymentStrategy) {
	c.paymentStrategy = strategy
}

func main() {
	creditCard := &CreditCardPayment{}
	context := PaymentContext{}
	context.setPaymentStrategy(creditCard)
	context.pay(10)
}
