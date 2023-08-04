package domain

import "finance/m/v2/domain/consts/month"

type TimeProgressor interface {
	ProgressMonths(startMonth month.Month, monthsQtd int, savings *Savings)
}
