package domain

type Accountant interface {
	Apply(startValue float64, monthString string) float64
}
