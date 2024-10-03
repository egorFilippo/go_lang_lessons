package models

import (
	"time"
)

type SberPayment struct {
	Id        string
	Timestamp float64
	Sum       int
	Type      PaymentType
}

func (p *SberPayment) GetCreatedAt() time.Time {
	t := time.Unix(int64(p.Timestamp), 0)
	return t
}

func (p *SberPayment) GetAmount() int {
	return p.Sum
}

func (p *SberPayment) GetType() PaymentType {
	return p.Type
}
