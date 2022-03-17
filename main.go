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
	logger.Log("Application starting")

	var dateTimeOneString string
	flag.StringVar(&dateTimeOneString, "start", "02/01/2011", "This time check by syetem")
	var dateTimeTwoString string
	flag.StringVar(&dateTimeTwoString, "end", "15/01/2011", "This time check by syetem")
	flag.Parse()

	if dateTimeOneString == "" || dateTimeTwoString == "" {
		logger.LogError(i18n.ShowText("NEED_TWO_DATE_STRING"))
		os.Exit(1)
	}

	dateTimeOne, err := date.DateBuilder(dateTimeOneString)
	if err != nil {
		logger.LogError(err.Error())
		os.Exit(1)
	}

	dateTimeTwo, err := date.DateBuilder(dateTimeTwoString)
	if err != nil {
		logger.LogError(err.Error())
		os.Exit(1)
	}

	distance, err := date.Distance(dateTimeTwo, dateTimeOne)
	if err != nil {
		logger.LogError(err.Error())
		os.Exit(1)
	}

	logger.Log(fmt.Sprintf("Result is %d", distance))
	logger.Log("Application stoping without error")
}
