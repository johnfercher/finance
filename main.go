package main

import (
	"finance/m/v2/contract"
	"finance/m/v2/domain"
	"finance/m/v2/domain/consts/month"
	"finance/m/v2/services"
	"fmt"
)

func main() {
	parameters, err := contract.LoadParameters()
	if err != nil {
		panic(err)
	}

	interestConverter := services.NewInterestConverter()

	yearCdiPercent := parameters.YearCDIPercent
	monthCdi := interestConverter.AnnualToMonth(yearCdiPercent)
	fmt.Printf("CDI: %f\n", monthCdi)

	currentMonth := parameters.Month
	monthsDuration := parameters.MonthsDuration

	bank := parameters.Savings.Bank
	cashback := parameters.Savings.Cashback
	fgts := parameters.Savings.FGTS

	gains := domain.NewGains()
	for _, taxable := range parameters.Gains.Taxables {
		gains.AddTaxableGain(taxable.Key, taxable.Value)
	}

	for _, nonTaxable := range parameters.Gains.NonTaxables {
		gains.AddNonTaxableGain(nonTaxable.Key, nonTaxable.Value)
	}

	gains.Print()

	spents := domain.NewSpents()

	for _, debit := range parameters.Spents.Debits {
		spents.AddDebitValue(debit.Key, debit.Value)
	}

	for _, credit := range parameters.Spents.Credits {
		spents.AddCreditValue(credit.Key, credit.Value)
	}

	spents.AddRemainingCreditValueToReachTotal(parameters.Spents.CreditTotal)

	spents.Print()

	balance := domain.NewBalance(gains, spents)
	fmt.Printf("Balance: %f\n", balance.Get())

	fgtsAccountant := services.NewFgtsAccountant(gains)
	salaryAccountant := services.NewSalaryAccountant(balance, monthCdi)
	cashbackAccountant := services.NewCashbackAccontant(spents, monthCdi)

	fmt.Printf("bank: %.2f, cashback: %.2f, fgts: %.2f, sum: %.2f, START\n", bank, cashback, fgts, bank+cashback+fgts)

	year := 0
	for i := 0; i < monthsDuration; i++ {
		fgts, _ = fgtsAccountant.Apply(fgts, currentMonth)
		newBank, bankRevenue := salaryAccountant.Apply(bank, currentMonth)
		newCashback, cashbackRevenue := cashbackAccountant.Apply(cashback, currentMonth)
		bank = newBank
		cashback = newCashback

		fmt.Printf("bank: %.2f, bankRevenue: %.2f, cashback: %.2f, cashbackRecenue: %.2f, fgts: %.2f, sum: %.2f, %s\n", bank, bankRevenue, cashback, cashbackRevenue, fgts, bank+cashback+fgts, currentMonth)

		currentMonth = month.GetNextMonth(currentMonth)
		if currentMonth == month.January {
			year++
			fmt.Printf("Year: %d\n", year)
		}
	}
}
