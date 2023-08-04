package entities

import (
	"finance/m/v2/domain/consts"
	"finance/m/v2/domain/consts/gain"
	"finance/m/v2/domain/consts/month"
)

type Gain struct {
	Label      string
	Value      float64
	Type       gain.Type
	Recurrence map[month.Month]bool
}

func NewGain(label string, value float64, gainType gain.Type, recurrence []month.Month) *Gain {
	return &Gain{
		Label:      label,
		Value:      value,
		Type:       gainType,
		Recurrence: month.BuildRecurrence(recurrence),
	}
}

type Gains struct {
	taxable    []*Gain
	nonTaxable []*Gain
}

func NewGains() *Gains {
	return &Gains{}
}

func (g *Gains) Print() {
	//fmt.Printf("TaxableBruteValue: %.2f, NonTaxableValue: %.2f, FGTS: %.2f, Liquid: %.2f, Total Liquid: %.2f\n", g.GetTaxableTotal(), g.GetNonTaxableTotal(), g.GetFGTS(), g.GetTaxableLiquid(), g.GetTotalLiquid())
}

func (g *Gains) AddTaxableGain(label string, value float64, recurrence []month.Month) *Gains {
	g.taxable = append(g.taxable, NewGain(label, value, gain.Taxable, recurrence))
	return g
}

func (g *Gains) AddNonTaxableGain(label string, value float64, recurrence []month.Month) *Gains {
	g.nonTaxable = append(g.nonTaxable, NewGain(label, value, gain.NonTaxable, recurrence))
	return g
}

func (g *Gains) GetTaxableTotal(month month.Month) float64 {
	total := 0.0
	for _, taxable := range g.taxable {
		if taxable.Recurrence[month] {
			total += taxable.Value
		}
	}

	return total
}

func (g *Gains) GetNonTaxableTotal(month month.Month) float64 {
	total := 0.0
	for _, nonTaxable := range g.nonTaxable {
		if nonTaxable.Recurrence[month] {
			total += nonTaxable.Value
		}
	}

	return total
}

func (g *Gains) GetTaxableLiquid(month month.Month) float64 {
	return g.GetTaxableTotal(month) * (1.0 - consts.SalaryTaxPercent)
}

func (g *Gains) GetFGTS(month month.Month) float64 {
	return g.GetTaxableTotal(month) * consts.FGTSPercent
}

func (g *Gains) GetTotalLiquid(month month.Month) float64 {
	return g.GetTaxableLiquid(month) + g.GetNonTaxableTotal(month)
}
