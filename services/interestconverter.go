package services

import (
	"finance/m/v2/domain"
	"finance/m/v2/domain/consts"
	"math"
)

type interestConverter struct {
}

func NewInterestConverter() domain.InterestConverter {
	return &interestConverter{}
}

func (i *interestConverter) AnnualToMonth(f float64) float64 {
	base := f/100.0 + 1.0
	return math.Pow(base, 1.0/float64(consts.Year)) - 1.0
}
