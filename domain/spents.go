package domain

import (
	"finance/m/v2/domain/consts/payment"
	"fmt"
)

type Spent struct {
	Label       string
	Value       float64
	PaymentType payment.Type
}

func NewSpent(label string, value float64, paymentType payment.Type) *Spent {
	return &Spent{
		Label:       label,
		Value:       value,
		PaymentType: paymentType,
	}
}

type Spents struct {
	Credit []*Spent
	Debit  []*Spent
}

func NewSpents() *Spents {
	return &Spents{}
}

func (s *Spents) Print() {
	fmt.Println("CREDIT")
	for _, credit := range s.Credit {
		fmt.Printf("%s: %.2f\n", credit.Label, credit.Value)
	}
	totalCredit := s.GetTotalCredit()
	fmt.Printf("Total Credit: %.2f\n", totalCredit)
	fmt.Println()

	fmt.Println("DEBIT")
	for _, debit := range s.Debit {
		fmt.Printf("%s: %.2f\n", debit.Label, debit.Value)
	}
	totalDebit := s.GetTotalDebit()
	fmt.Printf("Total Debit: %.2f\n", totalDebit)
	fmt.Println()
	fmt.Printf("TOTAL: %.2f\n", totalCredit+totalDebit)
	fmt.Println()
}

func (s *Spents) GetTotalCredit() float64 {
	total := 0.0
	for _, credit := range s.Credit {
		total += credit.Value
	}

	return total
}

func (s *Spents) GetTotalCreditCashback(percent float64) float64 {
	return s.GetTotalCredit() * percent
}

func (s *Spents) GetTotalDebit() float64 {
	total := 0.0
	for _, debit := range s.Debit {
		total += debit.Value
	}

	return total
}

func (s *Spents) GetTotal() float64 {
	return s.GetTotalCredit() + s.GetTotalDebit()
}

func (s *Spents) AddCreditValue(label string, value float64) *Spents {
	s.Credit = append(s.Credit, NewSpent(label, value, payment.CreditType))
	return s
}

func (s *Spents) AddRemainingCreditValueToReachTotal(value float64) *Spents {
	creditTotal := s.GetTotalCredit()
	if creditTotal > value {
		return s
	}

	s.Credit = append(s.Credit, NewSpent("other spents", value-creditTotal, payment.CreditType))
	return s
}

func (s *Spents) AddDebitValue(label string, value float64) *Spents {
	s.Debit = append(s.Debit, NewSpent(label, value, payment.DebitType))
	return s
}
