package util

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// isDoneDateValid (period string, doneDateString string) return bool
func isDoneDateValid(period string, doneDateString string) bool {
	// Split on forward comma.
	splitedStrings := strings.Split(doneDateString, "/")
	if len(splitedStrings) != 3 {
		return false
	}

	// Parse month/year of start/end
	startYearString := period[0:4]
	startMonthString := period[4:6]
	endYearString := period[7:11]
	endMonthString := period[11:13]
	startYear, _ := strconv.ParseInt(startYearString, 10, 64)
	startMonth, _ := strconv.ParseInt(startMonthString, 10, 64)
	endYear, _ := strconv.ParseInt(endYearString, 10, 64)
	endMonth, _ := strconv.ParseInt(endMonthString, 10, 64)

	// check month/year of done date in period
	doneYear, err := strconv.ParseInt(splitedStrings[0], 10, 64)
	isDoneDateValid := (err == nil) && (doneYear >= startYear) && (doneYear <= endYear)
	doneMonth, err := strconv.ParseInt(splitedStrings[1], 10, 64)
	isDoneDateValid = isDoneDateValid && (err == nil) && (doneMonth >= startMonth) && (doneMonth <= endMonth)
	return isDoneDateValid
}

// isPeriodValid (period string) return bool
func isPeriodValid(period string) bool {
	if len(period) != 13 {
		return false
	}
	isPeriodValid := true
	startYearString := period[0:4]
	startMonthString := period[4:6]
	endYearString := period[7:11]
	endMonthString := period[11:13]
	startYear, err1 := strconv.ParseInt(startYearString, 10, 64)
	startMonth, err2 := strconv.ParseInt(startMonthString, 10, 64)
	endYear, err3 := strconv.ParseInt(endYearString, 10, 64)
	endMonth, err4 := strconv.ParseInt(endMonthString, 10, 64)
	isPeriodValid = isPeriodValid && (endYear >= startYear)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return false
	}
	if endYear == startYear {
		isPeriodValid = isPeriodValid && (endMonth >= startMonth)
	}
	isPeriodValid = isPeriodValid && (startMonth >= 1) && (startMonth <= 12)
	isPeriodValid = isPeriodValid && (endMonth >= 1) && (endMonth <= 12)
	return isPeriodValid
}

// ArrangeData ...
// ###### [projectName] ######
// [2016/06]
func ArrangeData(reports [][]interface{}, period string) string {
	// Split period string : 201606-201612
	startYearString := period[0:4]
	starMonthString := period[4:6]
	endYearString := period[7:11]
	endMonthString := period[11:]
	// Convert to int
	startYear, _ := strconv.ParseInt(startYearString, 10, 64)
	startMonth, _ := strconv.ParseInt(starMonthString, 10, 64)
	endYear, _ := strconv.ParseInt(endYearString, 10, 64)
	endMonth, _ := strconv.ParseInt(endMonthString, 10, 64)
	// Result string buffer
	var stringBuffer bytes.Buffer
	// project list
	projects := getProjectsInReports(reports)
	// Foreach project
	for _, project := range projects {
		stringBuffer.WriteString("###### " + project + " ######\n")
		projectReports := getReportsByProject(reports, project)
		// output reports by monthly
		/**
		 * [2016/08]
		 * 1. aaa
		 * 2. bbb
		 */
		startDate := time.Date(int(startYear), time.Month(startMonth), 1, 0, 0, 0, 0, time.Local)
		endDate := time.Date(int(endYear), time.Month(endMonth), 2, 0, 0, 0, 0, time.Local)
		for startDate.Before(endDate.Local()) {
			// get monthly report
			year := startDate.Year()
			month := int(startDate.Month())
			monthReport := getRecordsByYearAndMonth(projectReports, year, month)
			// Print out if size > 0
			if len(monthReport) > 0 {
				// convert data to proper string
				yearPrint := fmt.Sprintf("%d", year)
				monthPrint := fmt.Sprintf("%02d", month)
				// put data to buffer
				stringBuffer.WriteString("[" + yearPrint + "/" + monthPrint + "]\n")
				for index, report := range monthReport {
					stringBuffer.WriteString(strconv.Itoa(index+1) + ". " + report[3].(string) + "\n")
				}
			}
			// Add 1 month to startDate
			startDate = startDate.AddDate(0, 1, 0)
		}
		stringBuffer.WriteString("\n")
	}
	return stringBuffer.String()
}

// get reports by year and month
func getRecordsByYearAndMonth(reports [][]interface{}, year int, month int) [][]interface{} {
	var wantedReports [][]interface{}
	for _, report := range reports {
		doneDate := report[5].(string)
		// get year, month from doneDate
		splitedString := strings.Split(doneDate, "/")
		doneYear, err1 := strconv.ParseInt(splitedString[0], 10, 64)
		doneMonth, err2 := strconv.ParseInt(splitedString[1], 10, 64)
		// error checking
		if err1 != nil && err2 != nil {
			continue
		}
		if int(doneYear) == year && int(doneMonth) == month {
			wantedReports = append(wantedReports, report)
		}
	}
	return wantedReports
}

// get reports by project
func getReportsByProject(reports [][]interface{}, projectName string) [][]interface{} {
	var wantedReports [][]interface{}
	for _, report := range reports {
		if report[2] == projectName {
			wantedReports = append(wantedReports, report)
		}
	}
	return wantedReports
}

// get projectName array in reports
func getProjectsInReports(reports [][]interface{}) []string {
	// project list
	var projects []string
	for _, record := range reports {
		// check if project name is valid
		projectString, ok := record[2].(string)
		if !ok {
			continue
		}
		// check if project is not in list
		if !isProjectInArray(projects, projectString) {
			projects = append(projects, projectString)
		}
	}
	return projects
}

// check if projectName is in given array
func isProjectInArray(projects []string, projectName string) bool {
	for _, project := range projects {
		if project == projectName {
			return true
		}
	}
	return false
}
