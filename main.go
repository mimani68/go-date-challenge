package main

import (
	"flag"
	"fmt"
	"os"

	"dev.io/v1/i18n"
	"dev.io/v1/pkg/date"
	"dev.io/v1/pkg/logger"
)

func main() {
	logger.Log(i18n.ShowText("APP_STARTING"))

	var startDateTimeString string
	flag.StringVar(&startDateTimeString, "start", "02/01/2011", "This time check by syetem")
	var endDateTimeString string
	flag.StringVar(&endDateTimeString, "end", "15/01/2011", "This time check by syetem")
	flag.Parse()

	if startDateTimeString == "" || endDateTimeString == "" {
		logger.LogError(i18n.ShowText("NEED_TWO_DATE_STRING"))
		os.Exit(1)
	}

	startDateTime, err := date.DateBuilder(startDateTimeString)
	if err != nil {
		os.Exit(1)
	}

	endDateTime, err := date.DateBuilder(endDateTimeString)
	if err != nil {
		os.Exit(1)
	}

	distance, err := date.Distance(endDateTime, startDateTime)
	if err != nil {
		logger.LogError(err.Error())
		os.Exit(1)
	}

	logger.Log(fmt.Sprintf("Result is %d", distance))
	logger.Log(i18n.ShowText("APP_STOP"))
}
