package implementing

import "fmt"

type PaymentProcessor interface {
	ProcessPayment(amount float64) string
}

type CreditCard struct{}

type Paypal struct{}

func (c CreditCard) ProcessPayment(amount float64) string {

	return fmt.Sprintf("Processing credit card payment %.2f", amount)
}

func (p Paypal) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Processing paypal payment %.2f", amount)
}

func ProcessPaymentFinal(pp PaymentProcessor, amount float64) string {
	return pp.ProcessPayment(amount)
}

func Task3() {
	var c CreditCard
	var p Paypal
	//c := PaymentProcessor //why can't we name like this ? :\
	result_c := ProcessPaymentFinal(c, 50.5)
	result_p := ProcessPaymentFinal(p, 6.3)

	fmt.Println(result_c)
	fmt.Println(result_p)

	// why it doesn't print result the way under
	//var payment1 PaymentProcessor
	//var payment2 PaymentProcessor
	//
	//payment1 = CreditCard{}
	//payment2 = Paypal{}
	//ProcessPaymentFinal(payment1, 100.5)
	//ProcessPaymentFinal(payment2, 10.5)

}
