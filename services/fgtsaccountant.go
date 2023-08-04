package services

import (
	"finance/m/v2/domain"
	"finance/m/v2/domain/consts/month"
	"finance/m/v2/domain/entities"
)

type fgtsAccountant struct {
	gains *entities.Gains
}

func NewFgtsAccountant(gains *entities.Gains) domain.Accountant {
	return &fgtsAccountant{
		gains: gains,
	}
}

func (t *fgtsAccountant) Apply(startValue float64, currentMonth month.Month) (float64, float64) {
	currentValue := startValue

	currentValue = currentValue + t.gains.GetFGTS(currentMonth)

	return currentValue, 0
}
