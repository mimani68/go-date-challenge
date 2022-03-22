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

func Parse(format string, value string) Time {
	time := strings.Split(value, "/")
	year, _ := strconv.Atoi(time[2])
	month, _ := strconv.Atoi(time[1])
	day, _ := strconv.Atoi(time[0])
	return Time{
		Year:   year,
		Month:  month,
		Day:    day,
		Hour:   0,
		Minute: 0,
		Second: 0,
	}
}
