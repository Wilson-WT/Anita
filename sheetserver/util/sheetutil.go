package util

import (
	"errors"
	"fmt"
	"os"

	setting "github.com/ibmboy19/Anita/sheetserver/config"
	errlog "github.com/inconshreveable/log15"
	sheets "google.golang.org/api/sheets/v4"
)

var eLog = errlog.New("package", "util")

// Append (dateTime string, user string, project string, content string, hours string, finishDate string)
func Append(dateTime, user, project, content, hours, finishDate string) int {
	client := RetrieveToken()
	srv, err := sheets.New(client)

	if err != nil {
		eLog.Error("Unable to retrieve Sheets Client %v", err)
		os.Exit(1)
	}
	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	spreadsheetID := setting.ImportConfig().GetGoogleSheetID()
	readRange := setting.ImportConfig().GetGoogleSheetTabName() + "!" + setting.ImportConfig().GetGoogleSheetRange()
	valueRange := &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values: [][]interface{}{
			[]interface{}{dateTime, user, project, content, hours, finishDate},
		},
	}
	resp, err := srv.Spreadsheets.Values.Append(spreadsheetID, readRange, valueRange).ValueInputOption("USER_ENTERED").InsertDataOption("INSERT_ROWS").Do()
	if err != nil {
		eLog.Error("Unable to write data to sheet. %v", err)
		os.Exit(1)
	} else {
		fmt.Println(resp.HTTPStatusCode)
	}
	return resp.HTTPStatusCode
}

// GetWorkingReports (accountName string, period string) return working reports array [][]
func GetWorkingReports(accountName, period string) ([][]interface{}, error) {
	if len(accountName) == 0 {
		return nil, errors.New("Account name is invalid")
	} else if !isPeriodValid(period) {
		return nil, errors.New("Period is invalid")
	}
	client := RetrieveToken()
	srv, err := sheets.New(client)
	if err != nil {
		eLog.Error("Unable to retrieve Sheets Client %v", err)
		os.Exit(1)
	}

	spreadsheetID := setting.ImportConfig().GetGoogleSheetID()
	readRange := setting.ImportConfig().GetGoogleSheetTabName() + "!" + setting.ImportConfig().GetGoogleSheetRange()
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		eLog.Error("Unable to retrieve data from sheet. %v", err)
		os.Exit(1)
	}
	var rowArray [][]interface{}
	for _, row := range resp.Values {
		// Print columns A and E, which correspond to indices 0 and 4.
		email := accountName + "@happygorgi.com"
		// row[1]: user's email
		// row[5]: done date of working report
		if row[1] == email && isDoneDateValid(period, row[5].(string)) {
			rowArray = append(rowArray, row)
		}
	}
	return rowArray, nil
}
