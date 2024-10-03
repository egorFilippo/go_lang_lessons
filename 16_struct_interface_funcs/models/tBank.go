package models

import "time"

type TBankPayment struct {
	Id       string
	Datetime string
	Amount   int
	Type     PaymentType
}

func (p *TBankPayment) GetCreatedAt() time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05Z", p.Datetime)
	return t
}

func (p *TBankPayment) GetAmount() int {
	return p.Amount
}

func (p *TBankPayment) GetType() PaymentType {
	return p.Type
}
