package services

import (
	"finance/m/v2/domain"
	"finance/m/v2/domain/consts"
	"finance/m/v2/domain/month"
)

type salaryAccountant struct {
	balance    *domain.Balance
	startMonth string
}

func NewSalaryAccountant(balance *domain.Balance) domain.Accountant {
	return &salaryAccountant{
		balance: balance,
	}
}

func (t *salaryAccountant) Apply(startValue float64, monthString string) float64 {
	currentValue := startValue

	if monthString == month.December {
		currentValue += t.balance.Gains.GetTotalLiquid()
	}

	currentValue = (currentValue + t.balance.Get()) * (1 + consts.CdiRate)
	return currentValue
}
