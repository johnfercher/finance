package domain

import (
	"finance/m/v2/domain/consts/month"
	"finance/m/v2/domain/entities"
)

type Balance struct {
	Gains  *entities.Gains
	Spents *entities.Spents
}

func NewBalance(gains *entities.Gains, spents *entities.Spents) *Balance {
	return &Balance{
		Gains:  gains,
		Spents: spents,
	}
}

func (b *Balance) Get(month month.Month) float64 {
	return b.Gains.GetTotalLiquid(month) - b.Spents.GetTotal(month)
}
