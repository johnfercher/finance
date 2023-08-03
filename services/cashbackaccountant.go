package services

import (
	"finance/m/v2/domain"
	"finance/m/v2/domain/consts"
)

type cashbackAccountant struct {
	spents *domain.Spents
	cdi    float64
}

func NewCashbackAccontant(spents *domain.Spents, cdi float64) domain.Accountant {
	return &cashbackAccountant{
		spents: spents,
		cdi:    cdi,
	}
}

func (t *cashbackAccountant) Apply(startValue float64, _ string) (float64, float64) {
	currentValue := startValue
	addedValue := currentValue + t.spents.GetTotalCreditCashback(consts.CashbackPercent)
	revenue := addedValue * (1 + t.cdi*2)

	return revenue, revenue - addedValue
}
