package main

import "payments/models"

func CalculateTotalSum(payments []Payment) (totalSum int) {
	for _, payment := range payments {
		totalSum += payment.GetAmount()
	}

	return
}

func CalculateAverageSum(payments []Payment) int {
	var totalSum int
	for _, payment := range payments {
		totalSum += payment.GetAmount()
	}
	return totalSum / len(payments)
}

func CalculatePaymentTotalIncome(payments []Payment) (totalSum int) {
	for _, payment := range payments {
		if payment.GetType() == models.Income {
			totalSum += payment.GetAmount()
		} else if payment.GetType() == models.Outcome {
			totalSum -= payment.GetAmount()
		}
	}
	return
}
