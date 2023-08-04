package services

import (
	"finance/m/v2/domain"
	"finance/m/v2/domain/consts"
	"finance/m/v2/domain/consts/month"
	"finance/m/v2/domain/entities"
)

type cashbackAccountant struct {
	spents *entities.Spents
	cdi    float64
}

func NewCashbackAccontant(spents *entities.Spents, cdi float64) domain.Accountant {
	return &cashbackAccountant{
		spents: spents,
		cdi:    cdi,
	}
}

func (t *cashbackAccountant) Apply(startValue float64, currentMonth month.Month) (float64, float64) {
	currentValue := startValue
	addedValue := currentValue + t.spents.GetTotalCreditCashback(currentMonth, consts.CashbackPercent)
	revenue := addedValue * (1 + t.cdi*2)

	return revenue, revenue - addedValue
}
