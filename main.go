package main

import (
	"flag"

	"dev.io/v1/pkg/date"
	"dev.io/v1/pkg/logger"
)

func main() {
	var dateTime string
	flag.StringVar(&dateTime, "date", "Something like 02/01/2011", "This time check by syetem")
	flag.Parse()
	logger.Log("Application starting")
	date, err := date.DateBuilder(dateTime)
	if err != nil {
		logger.LogError(err.Error())
	}
	logger.Log("Result is " + date.String())
	logger.Log("Application stoping")
}
