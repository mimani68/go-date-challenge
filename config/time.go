package config

import "time"

var (
	LOCATION         = time.UTC
	START_START_DATE = time.Date(1900, time.Month(1), 01, 0, 0, 1, 0, LOCATION)
	START_END_DATE   = time.Date(2999, time.Month(12), 31, 0, 0, 1, 0, LOCATION)
)
