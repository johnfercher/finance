package main

import (
	"finance/m/v2/contract"
	"finance/m/v2/domain"
	"finance/m/v2/domain/entities"
	"finance/m/v2/services"
	"fmt"
)

func main() {
	parameters, err := contract.LoadParameters()
	if err != nil {
		panic(err)
	}

	currentMonth := parameters.Month
	monthsDuration := parameters.MonthsDuration
	monthCdi := GetMonthCdi(parameters.YearCDIPercent)

	savings := &domain.Savings{
		Bank:     parameters.Savings.Bank,
		Cashback: parameters.Savings.Cashback,
		FGTS:     parameters.Savings.FGTS,
	}

	gains := BuildGains(parameters)
	spents := BuildSpents(parameters)
	balance := domain.NewBalance(gains, spents)

	fgtsAccountant := services.NewFgtsAccountant(gains)
	salaryAccountant := services.NewSalaryAccountant(balance, monthCdi)
	cashbackAccountant := services.NewCashbackAccontant(spents, monthCdi)

	passiveTimeProgressor := services.NewPassiveTimeProgressor(balance, fgtsAccountant, salaryAccountant, cashbackAccountant)

	fmt.Printf("bank: %.2f, cashback: %.2f, fgts: %.2f, sum: %.2f, START\n", savings.Bank, savings.Cashback, savings.FGTS, savings.Bank+savings.Cashback+savings.FGTS)
	passiveTimeProgressor.ProgressMonths(currentMonth, monthsDuration, savings)
}

func GetMonthCdi(yearCDIPercent float64) float64 {
	interestConverter := services.NewInterestConverter()
	return interestConverter.AnnualToMonth(yearCDIPercent)
}

func BuildGains(parameters *contract.Parameters) *entities.Gains {
	gains := entities.NewGains()
	for _, taxable := range parameters.Gains.Taxables {
		gains.AddTaxableGain(taxable.Label, taxable.Value, taxable.Recurrence)
	}

	for _, nonTaxable := range parameters.Gains.NonTaxables {
		gains.AddNonTaxableGain(nonTaxable.Label, nonTaxable.Value, nonTaxable.Recurrence)
	}

	return gains
}

func BuildSpents(parameters *contract.Parameters) *entities.Spents {
	spents := entities.NewSpents()

	for _, debit := range parameters.Spents.Debits {
		spents.AddDebitValue(debit.Label, debit.Value, debit.Recurrence)
	}

	for _, credit := range parameters.Spents.Credits {
		spents.AddCreditValue(credit.Label, credit.Value, credit.Recurrence)
	}

	spents.AddRemainingCreditValueToReachTotal(parameters.Spents.CreditTotal)

	spents.Print()

	return spents
}
