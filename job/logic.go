package job

import (
	"strconv"
	"time"
)

// calculate hour between 9:00~18:00
// exclude Saturday, Sunday
func calculateWorkingHour(startDate time.Time, endDate time.Time) string {
	const WorkingHour = 9
	startYear, startMon, startDay := startDate.Date()
	endYear, endMon, endDay := endDate.Date()

	var totalWorkingHoursString string
	if startYear == endYear && startMon == endMon && startDay == endDay {
		totalWorkingHours := endDate.Sub(startDate).Minutes() / 60
		totalWorkingHoursString = strconv.FormatFloat(totalWorkingHours, 'g', 2, 64)
	} else {
		// working hours of firstday
		firstDayEndTime := time.Date(startYear, startMon, startDay, 18, 0, 0, 0, time.Local)
		firstDayWorkingHours := firstDayEndTime.Sub(startDate).Minutes() / 60

		// working hours of lastday
		lastDayBeginingTime := time.Date(endYear, endMon, endDay, 9, 0, 0, 0, time.Local)
		lastDayWorkingHours := endDate.Sub(lastDayBeginingTime).Minutes() / 60

		// convert duration to workingdays
		duration := int(endDate.Sub(startDate).Hours() / 24)
		var workingDays = 0
		if endDate.Hour() >= startDate.Hour() {
			workingDays = duration + 1
		} else {
			workingDays = duration + 2
		}

		// calculate number of weekends
		startDate = time.Date(startYear, startMon, startDay, 0, 0, 0, 0, time.Local)
		endDate = time.Date(endYear, endMon, endDay, 0, 0, 0, 0, time.Local)
		numOfWeekends := 0
		for {
			if startDate.Equal(endDate) {
				break
			}
			if startDate.Weekday() == time.Saturday || startDate.Weekday() == time.Sunday {
				numOfWeekends++
			}
			startDate = startDate.Add(time.Hour * 24)
		}

		totalWorkingHours := firstDayWorkingHours + lastDayWorkingHours + float64((workingDays-2)*WorkingHour) - float64(numOfWeekends*WorkingHour)

		if (totalWorkingHours - float64(int(totalWorkingHours))) > 0 {
			totalWorkingHoursString = strconv.FormatFloat(totalWorkingHours, 'f', 1, 64)
		} else {
			totalWorkingHoursString = strconv.FormatInt(int64(totalWorkingHours), 10)
		}
	}
	return totalWorkingHoursString
}
