package date

import (
	"errors"
	"time"

	"dev.io/v1/i18n"
	"dev.io/v1/pkg/logger"
)

const (
	ISO_DATE_FORMAT     = "2006-01-02"
	ISO_DATETIME_FORMAT = "2006-01-02T12:20:11"
	ISO_CUSTOM_FORMAT   = "31/05/2006"
)

func DateBuilder(dateTime string) (time.Time, error) {
	dateOneObject, err := time.Parse(ISO_CUSTOM_FORMAT, dateTime)
	// dateOneObject, err := time.Parse(ISO_CUSTOM_FORMAT, dateTime)

	if err != nil {
		logger.Log(err.Error())
		return time.Time{}, errors.New(i18n.ShowText("INVALID_DATE"))
	}

	if !validator(dateOneObject) {
		logger.Log(i18n.ShowText("OUT_OF_RANGE"))
	}

	return dateOneObject, nil
}
