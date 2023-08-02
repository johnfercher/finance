package month

const (
	January   string = "January"
	February  string = "February"
	March     string = "March"
	April     string = "April"
	May       string = "May"
	June      string = "June"
	July      string = "July"
	August    string = "August"
	September string = "September"
	October   string = "October"
	November  string = "November"
	December  string = "December"
)

var monthNumber = map[string]int{
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

var numberMonth = map[int]string{
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

func GetMonthString(num int) string {
	return numberMonth[num]
}

func GetMonthNumber(month string) int {
	return monthNumber[month]
}

func GetNextMonth(month string) string {
	number := monthNumber[month]
	if number == 12 {
		number = 1
	} else {
		number++
	}

	return numberMonth[number]
}
