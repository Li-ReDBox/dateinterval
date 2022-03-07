package dateinterval

import (
	"errors"
	"strconv"
)

var ErrInvalidateDate = errors.New("invlidate date, it has to be in the format of dd/mm/yyyy between 01/01/1900 - 31/12/2999, is an allowed day in the month")

// validate checks if digits are valid days, months and years
func validate(parts []string) ([]int, error) {
	// parts are digits
	if len(parts) != 3 {
		return nil, ErrInvalidateDate
	}

	var d, m, y int

	if d, _ = strconv.Atoi(parts[0]); d < 1 || d > 31 {
		return nil, ErrInvalidateDate
	}
	if m, _ = strconv.Atoi(parts[1]); m < 1 || m > 12 {
		return nil, ErrInvalidateDate
	}
	if y, _ = strconv.Atoi(parts[2]); y < 1900 || y > 2999 {
		return nil, ErrInvalidateDate
	}
	if !isInMonth(d, m, y) {
		return nil, ErrInvalidateDate
	}
	return []int{d, m, y}, nil
}

// isInMonth checks if a day is in the month's range
func isInMonth(d, m, y int) bool {
	if LeapYear(y) && m == 2 && d <= 29 {
		return true
	}
	if d <= DAYS[m-1] {
		return true
	}
	return false
}
