package domain

import "finance/m/v2/domain/consts/month"

type Accountant interface {
	Apply(startValue float64, currentMonth month.Month) (float64, float64)
}
