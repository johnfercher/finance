package main

import (
	"finance/m/v2/contract"
	"finance/m/v2/domain"
	"finance/m/v2/domain/month"
	"finance/m/v2/services"
	"fmt"
)

func main() {
	parameters, err := contract.LoadParameters()
	if err != nil {
		panic(err)
	}

	currentMonth := parameters.Month

	savings := parameters.Savings.Bank
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

	fgtsAccountant := services.NewFgtsAccountant(gains)
	salaryAccountant := services.NewSalaryAccountant(balance)
	cashbackAccountant := services.NewCashbackAccontant(spents)

	fmt.Printf("savings: %.2f, cashback: %.2f, fgts: %.2f, sum: %.2f, START\n", savings, cashback, fgts, savings+cashback+fgts)

	for i := 0; i < 36; i++ {
		fgts = fgtsAccountant.Apply(fgts, currentMonth)
		savings = salaryAccountant.Apply(savings, currentMonth)
		cashback = cashbackAccountant.Apply(cashback, currentMonth)
		fmt.Printf("savings: %.2f, cashback: %.2f, fgts: %.2f, sum: %.2f, %s\n", savings, cashback, fgts, savings+cashback+fgts, currentMonth)

		currentMonth = month.GetNextMonth(currentMonth)
		if currentMonth == month.January {
			fmt.Println("New Year")
		}
	}
}
