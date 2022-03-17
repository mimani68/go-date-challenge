package date

import (
	"errors"
	"strings"
	"time"

	"dev.io/v1/i18n"
	"dev.io/v1/pkg/logger"
)

const (
	ISO_DATE_FORMAT     = "2006-01-02"
	ISO_DATETIME_FORMAT = time.RFC3339
)

func DateBuilder(dateTime string) (time.Time, error) {
	dateTimeSplitedArray := strings.Split(dateTime, "/")
	dateString := dateTimeSplitedArray[2] + "-" + dateTimeSplitedArray[1] + "-" + dateTimeSplitedArray[0] + "T00:00:00.000Z"
	dateObject, err := time.Parse(ISO_DATETIME_FORMAT, dateString)

	if err != nil {
		logger.LogError(err.Error())
		return time.Time{}, errors.New(i18n.ShowText("INVALID_DATE"))
	}

	if !validator(dateObject) {
		logger.LogError(i18n.ShowText("OUT_OF_RANGE"))
	}

	return dateObject, nil
}
