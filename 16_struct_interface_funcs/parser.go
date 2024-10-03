package main

import (
	"encoding/json"
	"io"
	"os"
	"payments/models"
)

func ParsePayment(filePath string) (payments []Payment, err error) {
	file, _ := os.Open(filePath)
	defer func() { err = file.Close() }()

	bytePaymentsData, _ := io.ReadAll(file)
	var rawPayments []interface{}
	err = json.Unmarshal(bytePaymentsData, &rawPayments)
	if err != nil {
		return nil, err
	}

	for _, rawPayment := range rawPayments {
		rawPaymentAsMap := rawPayment.(map[string]interface{})
		if rawPaymentAsMap["bank"] == "T-Bank" {
			payments = append(
				payments,
				&models.TBankPayment{
					Id:       rawPaymentAsMap["id"].(string),
					Datetime: rawPaymentAsMap["datetime"].(string),
					Amount:   int(rawPaymentAsMap["amount"].(float64)),
					Type:     models.PaymentType(rawPaymentAsMap["type"].(string)),
				},
			)
		} else {
			payments = append(
				payments,
				&models.SberPayment{
					Id:        rawPaymentAsMap["id"].(string),
					Timestamp: rawPaymentAsMap["timestamp"].(float64),
					Sum:       int(rawPaymentAsMap["sum"].(float64)),
					Type:      models.PaymentType(rawPaymentAsMap["type"].(string)),
				},
			)
		}
	}

	return
}
