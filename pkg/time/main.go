package time

import "fmt"

const (
	RFC3339 = "ISO"
	UTC     = "UTC"
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
	t.unixTime = (y*365 + m*29 + d) * 24 * 3600
	return t.unixTime
}

func (t *Time) String() string {
	return fmt.Sprintf("%d-%d-%dT00:00:00", t.Year, t.Month, t.Day)
}
