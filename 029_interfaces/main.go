package main

import "fmt"

//interface contract
type paymenter interface {
	pay(amount float32) // this is a method should implemented by all the objects
}

type Payment struct {
	gateway paymenter
}

//This function is tightly coupled with razorpay and stripe
//if we want to add a new payment gateway, we have to modify this function
//this voilets open closed principle of solid principles -> means open to extension but closed for modification
func (p Payment) makePayment(amount float32) {
	//razorpay payment gateway
	// razorpayPaymentGateway := razorpay{}
	// razorpayPaymentGateway.pay(amount)

	//stripe payment gateway
	// stripePaymentGateway := Stripe{}
	// stripePaymentGateway.pay(amount)

	p.gateway.pay(amount)
}

type razorpay struct {
}

func (r razorpay) pay(amount float32) {
	//logic to make payment
	fmt.Println("making payment using razor pay", amount)
}

type Stripe struct {
}

func (s Stripe) pay(amount float32) {
	fmt.Println("making payment using stripe", amount)
}

func main() {
	// stripePaymentGateway := Stripe{}
	razorpayPaymentGateway := razorpay{}
	newPayment := Payment{
		gateway: razorpayPaymentGateway,
	}
	newPayment.makePayment(100)
}
