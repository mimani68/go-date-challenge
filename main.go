package main

import (
	"flag"
	"fmt"

	"dev.io/v1/i18n"
	"dev.io/v1/pkg/date"
	"dev.io/v1/pkg/graceful_shutdown"
	"dev.io/v1/pkg/logger"
)

func main() {
	logger.Log(i18n.ShowText("APP_STARTING"))

	var startDateTimeString string
	flag.StringVar(&startDateTimeString, "start", "", "This time check by syetem")
	var endDateTimeString string
	flag.StringVar(&endDateTimeString, "end", "", "This time check by syetem")
	flag.Parse()

	if startDateTimeString == "" || endDateTimeString == "" {
		logger.LogError(i18n.ShowText("NEED_TWO_DATE_STRING"))
		graceful_shutdown.StopApplication(1, i18n.ShowText("STOP_WITH_ERROR"))
	}

	startDateTime, err := date.DateBuilder(startDateTimeString)
	if err != nil {
		graceful_shutdown.StopApplication(1, i18n.ShowText("STOP_WITH_ERROR"))
	}

	endDateTime, err := date.DateBuilder(endDateTimeString)
	if err != nil {
		graceful_shutdown.StopApplication(1, i18n.ShowText("STOP_WITH_ERROR"))
	}

	distance, err := date.Distance(endDateTime, startDateTime)
	if err != nil {
		logger.LogError(err.Error())
		graceful_shutdown.StopApplication(1, i18n.ShowText("STOP_WITH_ERROR"))
	}

	logger.Log(fmt.Sprintf("Result is %d", distance))
	graceful_shutdown.StopApplication(0, i18n.ShowText("APP_STOP"))
}
