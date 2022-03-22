package main

import (
	"fmt"

	time "dev.io/v1/pkg/time"
)

func main() {

	devSpace()

	// logger.Log(i18n.ShowText("APP_STARTING"))

	// var startDateTimeString string
	// flag.StringVar(&startDateTimeString, "start", "", "This time check by syetem")
	// var endDateTimeString string
	// flag.StringVar(&endDateTimeString, "end", "", "This time check by syetem")
	// flag.Parse()

	// if startDateTimeString == "" || endDateTimeString == "" {
	// 	logger.LogError(i18n.ShowText("NEED_TWO_DATE_STRING"))
	// 	graceful_shutdown.StopApplication(1, i18n.ShowText("STOP_WITH_ERROR"))
	// }

	// startDateTime, err := date.DateBuilder(startDateTimeString)
	// if err != nil {
	// 	graceful_shutdown.StopApplication(1, i18n.ShowText("STOP_WITH_ERROR"))
	// }

	// endDateTime, err := date.DateBuilder(endDateTimeString)
	// if err != nil {
	// 	graceful_shutdown.StopApplication(1, i18n.ShowText("STOP_WITH_ERROR"))
	// }

	// distance, err := date.Distance(endDateTime, startDateTime)
	// if err != nil {
	// 	logger.LogError(err.Error())
	// 	graceful_shutdown.StopApplication(1, i18n.ShowText("STOP_WITH_ERROR"))
	// }

	// logger.Log(fmt.Sprintf("Result is %d", distance))
	// graceful_shutdown.StopApplication(0, i18n.ShowText("APP_STOP"))
}

func devSpace() {
	// t := time.Now()  // FIXME: not for now plase ;)
	t := time.Date(2020, 12, 02, 12, 01, 01)
	// t := time.Parse(time.RFC3339, "03/01/2001")
	// t := time.Parse(time.RFC3339, "03/22/2022")
	fmt.Println(t.String())
	// fmt.Println(t.Unix())
	// fmt.Println(t.Before(time.Date(2007, 02, 02, 00, 01, 01)))
	// fmt.Println(t.After(time.Date(2007, 02, 02, 00, 01, 01)))
	// distance, _ := date.Distance(time.Date(2001, 01, 03, 00, 00, 00), time.Date(2001, 01, 01, 00, 00, 00))
	// distance, _ := date.Distance(time.Parse(time.RFC3339, "03/01/2001"), time.Parse(time.RFC3339, "01/01/2001"))
	// fmt.Println(distance)
}
