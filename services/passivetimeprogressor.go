package services

import (
	"finance/m/v2/domain"
	"finance/m/v2/domain/consts"
	"finance/m/v2/domain/consts/month"
	"fmt"
)

type passiveTimeProgressor struct {
	balance            *domain.Balance
	fgtsAccountant     domain.Accountant
	salaryAccountant   domain.Accountant
	cashbackAccountant domain.Accountant
}

func NewPassiveTimeProgressor(balance *domain.Balance, fgtsAccountant domain.Accountant, salaryAccountant domain.Accountant, cashbackAccountant domain.Accountant) *passiveTimeProgressor {
	return &passiveTimeProgressor{
		balance:            balance,
		fgtsAccountant:     fgtsAccountant,
		salaryAccountant:   salaryAccountant,
		cashbackAccountant: cashbackAccountant,
	}
}

func (t *passiveTimeProgressor) ProgressMonths(startMonth month.Month, monthsQtd int, savings *domain.Savings) {
	year := 0
	fgts := savings.FGTS
	cashback := savings.Cashback
	bank := savings.Bank

	currentMonth := startMonth
	for i := 0; i < monthsQtd; i++ {
		fgts, _ = t.fgtsAccountant.Apply(fgts, currentMonth)
		newBank, bankRevenue := t.salaryAccountant.Apply(bank, currentMonth)
		newCashback, cashbackRevenue := t.cashbackAccountant.Apply(cashback, currentMonth)
		bank = newBank
		cashback = newCashback

		fmt.Printf("bruteTaxable: %.2f, liquid: %.2f, spents: %.2f, balance: %.2f, cashbackGain: %.2f, fgtsGain: %.2f, bank: %.2f, bankRevenue: %.2f, cashback: %.2f, cashbackRecenue: %.2f, fgts: %.2f, sum: %.2f, %s\n", t.balance.Gains.GetTaxableTotal(currentMonth), t.balance.Gains.GetTotalLiquid(currentMonth), t.balance.Spents.GetTotal(currentMonth), t.balance.Get(currentMonth), t.balance.Spents.GetTotalCreditCashback(currentMonth, consts.CashbackPercent), t.balance.Gains.GetFGTS(currentMonth), bank, bankRevenue, cashback, cashbackRevenue, fgts, bank+cashback+fgts, currentMonth)

		currentMonth = month.GetNextMonth(currentMonth)
		if currentMonth == month.January {
			year++
			fmt.Printf("Year: %d\n", year)
		}
	}
}
