// interfcaes in go lang
package main

import "fmt"

// interface is a contract
// making interfaces -> there is the convention of naming interfaces as we have to put "er" on the last
type paymenter interface {
	pay(amount float32) //we have to put one method in interfaces
	refund(amount float32, accountID string)
}

// sceniro intergating payment gateway
type payment struct {
	gateway      paymenter
	refundAmount paymenter
}

// Open Close principle we are voiliting
func (p payment) makePayment(amount float32, accountID string) {
	//code modify krne ki need padi hme
	// razorpayPayment := razorpay{}
	// razorpayPayment.pay(amount)

	// stripePayment := stripe{}
	// stripePayment.pay(amount)

	p.gateway.pay(amount)
	p.refundAmount.refund(amount, accountID)
}

// NOTE: interface mei jo method signature hai woh same hoga tb go automatically implement ho jata hai.
type razorpay struct{}

func (r razorpay) pay(amount float32) {
	//logic to make payment
	fmt.Println("making payement using razorpay:", amount)
}

func (r razorpay) refund(amount float32, accountID string) {
	fmt.Printf("The refdung amount: %.2f is credifited to account %s\n", amount, accountID)
}

// if we want to implement stripe then how can we do
type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("making payment using stripe:", amount)
}

// for testign purpose
type fakePayment struct{}

func (f fakePayment) pay(amount float32) {
	fmt.Println("Making payment by third party gateways:", amount)
}

func main() {

	razorPaymentGW := razorpay{}
	// stripePaymentGW := stripe{}
	// fakePaymentGW := fakePayment{}

	newPay := payment{
		gateway:      razorPaymentGW,
		refundAmount: razorPaymentGW,
	}
	newPay.makePayment(90, "abc234")
}
