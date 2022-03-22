package time

import (
	"strconv"
	"strings"
)

func Now() Time {
	return Time{}
}

func Date(year, month, day, hour, minute, sec int) Time {
	return Time{
		Year:   year,
		Month:  month,
		Day:    day,
		Hour:   hour,
		Minute: minute,
		Second: sec,
	}
}

func Parse(format string, value string) (Time, error) {
	if format == RFC3339 {
		input := strings.Split(value, "T")
		time := strings.Split(input[0], "-")
		year, errYear := strconv.Atoi(time[0])
		if errYear != nil {
			return Time{}, errYear
		}
		month, errMonth := strconv.Atoi(time[1])
		if errMonth != nil {
			return Time{}, errMonth
		}
		day, erroDay := strconv.Atoi(time[2])
		if erroDay != nil {
			return Time{}, erroDay
		}
		return Time{
			Year:   year,
			Month:  month,
			Day:    day,
			Hour:   0,
			Minute: 0,
			Second: 0,
		}, nil
	} else {
		time := strings.Split(value, "/")
		year, errYear := strconv.Atoi(time[2])
		if errYear != nil {
			return Time{}, errYear
		}
		month, errMonth := strconv.Atoi(time[1])
		if errMonth != nil {
			return Time{}, errMonth
		}
		day, erroDay := strconv.Atoi(time[0])
		if erroDay != nil {
			return Time{}, erroDay
		}
		return Time{
			Year:   year,
			Month:  month,
			Day:    day,
			Hour:   0,
			Minute: 0,
			Second: 0,
		}, nil
	}
}
