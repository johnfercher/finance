package month

type Month string
type Recurrence map[Month]bool

const (
	January   Month = "January"
	February  Month = "February"
	March     Month = "March"
	April     Month = "April"
	May       Month = "May"
	June      Month = "June"
	July      Month = "July"
	August    Month = "August"
	September Month = "September"
	October   Month = "October"
	November  Month = "November"
	December  Month = "December"
)

var monthNumber = map[Month]int{
	January:   1,
	February:  2,
	March:     3,
	April:     4,
	May:       5,
	June:      6,
	July:      7,
	August:    8,
	September: 9,
	October:   10,
	November:  11,
	December:  12,
}

var numberMonth = map[int]Month{
	1:  January,
	2:  February,
	3:  March,
	4:  April,
	5:  May,
	6:  June,
	7:  July,
	8:  August,
	9:  September,
	10: October,
	11: November,
	12: December,
}

func GetNextMonth(month Month) Month {
	number := monthNumber[month]
	if number == 12 {
		number = 1
	} else {
		number++
	}

	return numberMonth[number]
}

func BuildRecurrence(arr []Month) map[Month]bool {
	recurrence := make(map[Month]bool)
	for _, value := range arr {
		recurrence[value] = true
	}

	return recurrence
}

func GetAnnualRecurrence() []Month {
	return []Month{
		January,
		February,
		March,
		April,
		May,
		June,
		July,
		August,
		September,
		October,
		November,
		December,
	}
}
