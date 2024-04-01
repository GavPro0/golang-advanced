package main

import "fmt"

type Payment interface {
	pay()
}

type CashPayment struct{}

func (CashPayment) pay() {
	fmt.Println("Payment using Cash")
}

func ProcessPayment(p Payment) {
	p.pay()
}

type BankPayment struct{}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Paying using BankAcount %d\n", bankAccount)
}

type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

func (bpa *BankPaymentAdapter) pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)

	// bank := &BankPayment{}
	// ProcessPayment(bank)

	bpa := &BankPaymentAdapter{
		bankAccount: 5,
		BankPayment: &BankPayment{},
	}
	ProcessPayment(bpa)

}
