package date

import (
	"time"

	"dev.io/v1/config"
)

func validator(dateTime time.Time) bool {
	return dateTime.After(config.TIME_START) && dateTime.Before(config.TIME_END)
}
