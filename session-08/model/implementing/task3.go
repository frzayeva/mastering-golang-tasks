package implementing

import "fmt"

type PaymentProcessor interface {
	ProcessPayment(amount float64) float64
}

type CreditCard struct {
	amount float64
}

type Paypal struct {
	amount float64
}

func (c CreditCard) ProcessPayment(amount float64) float64 {

	return c.amount
}

func (p Paypal) ProcessPayment(amount float64) float64 {
	return p.amount
}

func Task3() {
	var c PaymentProcessor
	var p PaymentProcessor
	//c := PaymentProcessor //why can't we name like this ? :\
	c = CreditCard{
		amount: 30.5,
	}

	p = Paypal{
		amount: 20.5,
	}
	fmt.Println(c.ProcessPayment(3.5))
	fmt.Println(p.ProcessPayment(3.5))

}
