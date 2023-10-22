package utils

import (
	"strconv"
	"time"
)

func DateToStdNow() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 15:04:05")
}

func StrToTimeStamp(param string) time.Time {
	dateNow, _ := time.Parse("2006-01-02", param)
	return dateNow
}

func TimeStampToDate(param time.Time) string {

	format := strconv.Itoa(param.Year()) + "-" + strconv.Itoa(int(param.Month())) + "-" + strconv.Itoa(param.Day())
	return format
}
