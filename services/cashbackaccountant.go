package services

import (
	"finance/m/v2/domain"
	"finance/m/v2/domain/consts"
)

type cashbackAccountant struct {
	spents *domain.Spents
}

func NewCashbackAccontant(spents *domain.Spents) domain.Accountant {
	return &cashbackAccountant{
		spents: spents,
	}
}

func (t *cashbackAccountant) Apply(startValue float64, monthString string) float64 {
	currentValue := startValue
	currentValue = (currentValue + t.spents.GetTotalCreditCashback(consts.CashbackPercent)) * (1 + consts.DoubleCdiRate)

	return currentValue
}
