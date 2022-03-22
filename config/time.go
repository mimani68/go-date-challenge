package config

import "dev.io/v1/pkg/time"

var (
	LOCATION   = time.UTC
	TIME_START = time.Date(1900, 1, 01, 0, 0, 1)
	TIME_END   = time.Date(2999, 12, 31, 0, 0, 1)
)
