package main

import (
	"dev.io/v1/i18n"
	"dev.io/v1/pkg/logger"
	"dev.io/v1/pkg/validator"
)

func main() {
	dateOne := "2011-12-10"
	if !validator.DateValidator(dateOne) {
		logger.Log(i18n.ShowText("INVALID_DATE"))
	}
	dateTwo := "2012-12-10"
	if !validator.DateValidator(dateTwo) {
		logger.Log(i18n.ShowText("INVALID_DATE"))
	}
	logger.Log("end of programm")
}
