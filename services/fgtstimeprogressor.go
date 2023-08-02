package services

import (
	"finance/m/v2/domain"
	"finance/m/v2/domain/month"
)

type fgtsAccountant struct {
	gains *domain.Gains
}

func NewFgtsAccountant(gains *domain.Gains) domain.Accountant {
	return &fgtsAccountant{
		gains: gains,
	}
}

func (t *fgtsAccountant) Apply(startValue float64, monthString string) float64 {
	currentValue := startValue

	if monthString == month.December {
		currentValue += t.gains.GetFGTS()
	}

	currentValue = currentValue + t.gains.GetFGTS()
	return currentValue
}
