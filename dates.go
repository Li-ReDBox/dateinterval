package dateinterval

import (
	"fmt"
	"regexp"
)

var DAYS = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// DaysSoFar calculates the number of days since January 1 to the date
func DaysSoFar(d, m int) int {
	ds := d
	for i := 0; i < m-1; i++ {
		ds = ds + DAYS[i]
	}
	return ds
}

// LeapYear checks if a given year is a leap year
func LeapYear(y int) bool {
	if y%4 == 0 {
		if y%100 != 0 {
			return true
		}
		if y%400 == 0 {
			return true
		}
	}
	return false
}

// LeapYears calculates the total number of leap years between two years
func LeapYears(ys, ye int) int {
	total := 0
	for i := ys; i <= ye; i++ {
		if LeapYear(i) {
			total += 1
		}
	}
	return total
}

// A Date reprents a date
type Date struct {
	Year, Month, Day int
}

// CreateDate returns a Date from the given string
func CreateDate(ds string) (Date, error) {
	d, err := parseDateNumbers(ds)
	if err != nil {
		return Date{}, err
	}
	return Date{d[2], d[1], d[0]}, nil
}

// String is Stringer of Date
func (d Date) String() string {
	return fmt.Sprintf("%02d/%02d/%4d", d.Day, d.Month, d.Year)
}

// toInts returns an int array of the current Date
func (d Date) toInts() []int {
	return []int{d.Day, d.Month, d.Year}
}

// Interval returns the number of days between to Dates
func (d Date) Interval(o Date) int {
	small := d
	big := o
	if d.Year > o.Year || (d.Year == o.Year && d.Month > o.Month) ||
		(d.Year == o.Year && d.Month == o.Month && d.Day > o.Day) {
		small = o
		big = d
	}

	return diff(small.toInts(), big.toInts())
}

// diff calculates the how many days are between two dates.
func diff(dateSmall, dateBig []int) int {
	ds := DaysSoFar(dateSmall[0], dateSmall[1])
	de := DaysSoFar(dateBig[0], dateBig[1])
	var delta int
	// same year
	if dateSmall[2] == dateBig[2] {
		// same dates, the Count is zero
		delta = de - ds
		if delta == 0 {
			return 0
		}
		// check if this involve Feb in a leap year
		if dateSmall[1] <= 2 && dateBig[1] > 2 && LeapYear(dateBig[2]) {
			return delta
		}
		return delta - 1
	}
	lyears := 0
	if dateBig[1] > 2 {
		lyears += LeapYears(dateBig[2], dateBig[2])
	}
	if dateSmall[1] < 2 {
		lyears += LeapYears(dateSmall[2], dateSmall[2])
	}
	if dateBig[2]-dateSmall[2] > 1 {
		lyears += LeapYears(dateSmall[2]+1, dateBig[2]-1)
	}
	ys := dateBig[2] - dateSmall[2]
	return ys*365 + lyears + de - ds - 1
}

// getParts generates a string array from a string in dd/mm/yyyy format
func getParts(dateStr string) []string {
	// if day or month is single digit, they cannot be zero
	re := regexp.MustCompile(`^(\d{1,2}?)/(\d{1,2}?)/(\d{4})$`)
	parts := re.FindStringSubmatch(dateStr)
	if parts != nil {
		return parts[1:]
	}
	return nil
}

// parseDateNumbers parses a stirng using dd/mm/yyyy format into an int array.
func parseDateNumbers(dateStr string) ([]int, error) {
	parts := getParts(dateStr)
	if parts == nil {
		return nil, ErrInvalidateDate
	}
	return validate(parts)
}
