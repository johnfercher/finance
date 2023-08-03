package domain

type InterestConverter interface {
	AnnualToMonth(float64) float64
}
