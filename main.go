package main

import (
	"finance/m/v2/domain"
	"finance/m/v2/domain/month"
	"finance/m/v2/services"
	"fmt"
)

func main() {
	savings := 10000.0
	fgts := 10000.0
	cashback := 200.98
	currentMonth := month.August

	gains := domain.NewGains().
		AddTaxableGain("salary", 14000.0).
		AddNonTaxableGain("VR/VA", 1000.0)

	gains.Print()

	spents := domain.NewSpents().
		AddDebitValue("aluguel", 1550).
		AddDebitValue("condominio", 0).
		AddDebitValue("IPTU", 28.58).
		AddDebitValue("internet", 39.95).
		AddDebitValue("health", 1636).
		AddDebitValue("energy", 125).
		AddCreditValue("spotify", 34.9).
		AddCreditValue("netflix", 55.9).
		AddCreditValue("meli 6", 19.46).
		AddCreditValue("tim beta", 60.0).
		AddCreditValue("hbo max", 13.95).
		AddCreditValue("remédios", 300).
		AddRemainingCreditValueToReachTotal(3500)

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
