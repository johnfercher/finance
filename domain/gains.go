package domain

import (
	"finance/m/v2/domain/consts"
	"finance/m/v2/domain/gain"
	"fmt"
)

type Gain struct {
	Label    string
	Value    float64
	GainType gain.GainType
}

func NewGain(label string, value float64, gainType gain.GainType) *Gain {
	return &Gain{
		Label:    label,
		Value:    value,
		GainType: gainType,
	}
}

type Gains struct {
	Taxable    []*Gain
	NonTaxable []*Gain
}

func NewGains() *Gains {
	return &Gains{}
}

func (g *Gains) Print() {
	fmt.Printf("TaxableBruteValue: %.2f, NonTaxableValue: %.2f, FGTS: %.2f, Liquid: %.2f, Total Liquid: %.2f\n", g.GetTaxableTotal(), g.GetNonTaxableTotal(), g.GetFGTS(), g.GetTaxableLiquid(), g.GetTotalLiquid())
}

func (g *Gains) AddTaxableGain(label string, value float64) *Gains {
	g.Taxable = append(g.Taxable, NewGain(label, value, gain.TaxableGainType))
	return g
}

func (g *Gains) AddNonTaxableGain(label string, value float64) *Gains {
	g.NonTaxable = append(g.NonTaxable, NewGain(label, value, gain.NonTaxableGainType))
	return g
}

func (g *Gains) GetTaxableTotal() float64 {
	total := 0.0
	for _, taxable := range g.Taxable {
		total += taxable.Value
	}

	return total
}

func (g *Gains) GetNonTaxableTotal() float64 {
	total := 0.0
	for _, nonTaxable := range g.NonTaxable {
		total += nonTaxable.Value
	}

	return total
}

func (g *Gains) GetTaxableLiquid() float64 {
	return g.GetTaxableTotal() * (1.0 - consts.SalaryTaxPercent)
}

func (g *Gains) GetFGTS() float64 {
	return g.GetTaxableTotal() * consts.FGTSPercent
}

func (g *Gains) GetTotalLiquid() float64 {
	return g.GetTaxableLiquid() + g.GetNonTaxableTotal()
}
