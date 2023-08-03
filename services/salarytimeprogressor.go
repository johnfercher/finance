package services

import (
	"finance/m/v2/domain"
	"finance/m/v2/domain/consts/month"
)

type salaryAccountant struct {
	balance    *domain.Balance
	cdi        float64
	startMonth string
}

func NewSalaryAccountant(balance *domain.Balance, cdi float64) domain.Accountant {
	return &salaryAccountant{
		balance: balance,
		cdi:     cdi,
	}
}

func (t *salaryAccountant) Apply(startValue float64, monthString string) (float64, float64) {
	currentValue := startValue

	if monthString == month.December {
		currentValue += t.balance.Gains.GetTotalLiquid()
	}

	addedValue := (currentValue + t.balance.Get())
	revenueValue := addedValue * (1 + t.cdi)

	return revenueValue, revenueValue - addedValue
}
