package date

import (
	"time"

	"dev.io/v1/config"
)

func validator(dateTime time.Time) bool {
	return dateTime.After(config.START_END_DATE) && dateTime.Before(config.START_END_DATE)
}
