package time

import (
	"fmt"
)

const (
	RFC3339 = "ISO"
	UTC     = "UTC"
	CUSTOM  = "*"
)

type Time struct {
	unixTime int
	Year     int
	Month    int
	Day      int
	Hour     int
	Minute   int
	Second   int
}

func (t *Time) Format(formatString string) string {
	resut := ""
	switch formatString {
	case RFC3339:
		resut = fmt.Sprintf("%d-%d-%dT00:00:00", t.Year, t.Month, t.Day)
	case UTC:
		resut = fmt.Sprintf("%d-%d-%dT00:00:00", t.Year, t.Month, t.Day)
	}
	return resut
}

func (t *Time) Before(time Time) bool {
	return t.unix() > time.unix()
}

func (t *Time) After(time Time) bool {
	return !t.Before(time)
}

func (t *Time) Unix() int {
	return t.unix()
}

func (t *Time) unix() int {
	y := (t.Year - 1970)
	m := t.Month
	d := t.Day
	t.unixTime = (y*365 + t.cumulativeDayForMonth(m) + d + t.cumulativeLeapYearDays(t.Year)) * 24 * 3600
	return t.unixTime
}

func (t *Time) String() string {
	return fmt.Sprintf("%d-%d-%dT00:00:00", t.Year, t.Month, t.Day)
}

func (t *Time) cumulativeLeapYearDays(year int) int {
	result := 0
	for i := 1970; i < year; i++ {
		if t.isLeapYear(i) {
			result = result + 1
		}
	}
	return result
}

func (t *Time) isLeapYear(year int) bool {
	if 0 == year%4 && 0 != year%100 || 0 == year%400 {
		return true
	} else {
		return false
	}
}

func (t *Time) cumulativeDayForMonth(month int) int {
	result := 0
	dayList := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	for i := 0; i < month; i++ {
		result = result + dayList[i]
	}
	return result
}
