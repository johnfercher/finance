package entities

import (
	"finance/m/v2/domain/consts/month"
	"finance/m/v2/domain/consts/payment"
	"fmt"
)

type Spent struct {
	Label       string
	Value       float64
	PaymentType payment.Type
	Recurrence  map[month.Month]bool
}

func NewSpent(label string, value float64, paymentType payment.Type, recurrence []month.Month) *Spent {
	return &Spent{
		Label:       label,
		Value:       value,
		PaymentType: paymentType,
		Recurrence:  month.BuildRecurrence(recurrence),
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
	/*fmt.Println("CREDIT")
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
	fmt.Println()*/
}

func (s *Spents) GetTotalCredit(month month.Month) float64 {
	total := 0.0
	for _, credit := range s.Credit {
		if credit.Recurrence[month] {
			total += credit.Value
		}
	}

	return total
}

func (s *Spents) GetTotalDebit(month month.Month) float64 {
	total := 0.0
	for _, debit := range s.Debit {
		if debit.Recurrence[month] {
			total += debit.Value
		}
	}

	return total
}

func (s *Spents) GetTotal(month month.Month) float64 {
	return s.GetTotalCredit(month) + s.GetTotalDebit(month)
}

func (s *Spents) AddCreditValue(label string, value float64, recurrence []month.Month) *Spents {
	s.Credit = append(s.Credit, NewSpent(label, value, payment.CreditType, recurrence))
	return s
}

func (s *Spents) AddRemainingCreditValueToReachTotal(value float64) *Spents {
	recurrence := month.GetAnnualRecurrence()
	for _, recurrenceMonth := range recurrence {
		creditTotal := s.GetTotalCredit(recurrenceMonth)
		if creditTotal > value {
			continue
		}

		s.Credit = append(s.Credit, NewSpent(fmt.Sprintf("other spent %s", recurrenceMonth), value-creditTotal, payment.CreditType, []month.Month{recurrenceMonth}))
	}

	return s
}

func (s *Spents) AddDebitValue(label string, value float64, recurrence []month.Month) *Spents {
	s.Debit = append(s.Debit, NewSpent(label, value, payment.DebitType, recurrence))
	return s
}

func (s *Spents) GetTotalCreditCashback(month month.Month, percent float64) float64 {
	return s.GetTotalCredit(month) * percent
}
