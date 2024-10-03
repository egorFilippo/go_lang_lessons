package main

import (
	"fmt"
	"os"
	"payments/models"
	"time"
)

type Payment interface {
	GetCreatedAt() time.Time
	GetAmount() int
	GetType() models.PaymentType
}

func testFunc(arg Payment) {
	arg.GetAmount()
}

func main() {
	sberPayment := models.TBankPayment{}
	tbankPayment := models.SberPayment{}
	testFunc(&sberPayment)
	testFunc(&tbankPayment)

	payments, err := ParsePayment("payments.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("CalculateTotalSum: ", CalculateTotalSum(payments))
	fmt.Println("CalculateAverageSum: ", CalculateAverageSum(payments))
	fmt.Println("CalculatePaymentTotalIncome: ", CalculatePaymentTotalIncome(payments))

}
