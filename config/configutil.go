package config

import (
	"strconv"
	"strings"
)

// GetCronSettingForArchiveReportedCards ...
func GetCronSettingForArchiveReportedCards(archiveReportedCardsWeekDay, autoArchiveTime string) string {
	splitedStrings := strings.Split(autoArchiveTime, ":")
	originString := splitedStrings[0]
	hour, _ := strconv.ParseInt(originString, 10, 64)
	if hour == 24 {
		hour = 0
	}
	newAutoArchiveTimeString := strconv.FormatInt(hour, 10)
	return "0 0 " + newAutoArchiveTimeString + " * * " + archiveReportedCardsWeekDay
}
