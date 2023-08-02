package domain

type Balance struct {
	Gains  *Gains
	Spents *Spents
}

func NewBalance(gains *Gains, spents *Spents) *Balance {
	return &Balance{
		Gains:  gains,
		Spents: spents,
	}
}

func (b *Balance) Get() float64 {
	return b.Gains.GetTotalLiquid() - b.Spents.GetTotal()
}
