package date

import (
	"dev.io/v1/config"
	"dev.io/v1/pkg/time"
)

func DateValidator(dateTime time.Time) bool {
	return dateTime.After(config.TIME_END) && dateTime.Before(config.TIME_START)
}
