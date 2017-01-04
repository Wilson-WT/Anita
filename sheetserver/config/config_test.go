package config

import (
	"io/ioutil"
	"log"
	"testing"
)

//// TestCheckAllConfig
func TestCheckAllConfig_Valid(t *testing.T) {
	configPath = "config_test.ini"

	// Prepare test data
	d1 := []byte("[Google Sheets API]\n" +
		"SheetID = testDB36KqKEFpWTn8DpuOtJM_vOpTMtys_RQeMAOJhI\n" +
		"TabName = 表單回應 1\n" +
		"Range = A2:G\n")
	err := ioutil.WriteFile(configPath, d1, 0644)
	if err != nil {
		log.Fatal("Unable to wirte data")
	}

	// check all config
	CheckAllConfig()
}

//// TestGetGoogleSheetID
func TestGetGoogleSheetID_Valid(t *testing.T) {
	configPath = "config_test.ini"
	// Prepare test data
	d1 := []byte("[Google Sheets API]\nSheetID = 123nDB36KqKEFpWTn8DpuOtJM_vOpTMtys_RQeMAOXyz")
	err := ioutil.WriteFile(configPath, d1, 0644)
	if err != nil {
		log.Fatal("Unable to wirte data")
	}
	// assert content
	GoogleSheetID := ImportConfig().GetGoogleSheetID()
	if GoogleSheetID == "123nDB36KqKEFpWTn8DpuOtJM_vOpTMtys_RQeMAOXyz" {
		t.Log("TestGetGoogleSheetID PASS")
	} else {
		t.Error("Content of GoogleSheetID is not correct")
	}
}

//// TestGetGoogleSheetTabName
func TestGetGoogleSheetTabName_Valid(t *testing.T) {
	configPath = "config_test.ini"
	// Prepare test data
	d1 := []byte("[Google Sheets API]\nTabName = qatest")
	err := ioutil.WriteFile(configPath, d1, 0644)
	if err != nil {
		log.Fatal("Unable to wirte data")
	}
	// assert content
	GoogleSheetTabName := ImportConfig().GetGoogleSheetTabName()
	if GoogleSheetTabName == "qatest" {
		t.Log("TestGetGoogleSheetTabName PASS")
	} else {
		t.Error("Content of GoogleSheetID is not correct")
	}
}

//// TestGetGoogleSheetRange
func TestGetGoogleSheetRange_Valid(t *testing.T) {
	configPath = "config_test.ini"
	// Prepare test data
	d1 := []byte("[Google Sheets API]\nRange = A2:F")
	err := ioutil.WriteFile(configPath, d1, 0644)
	if err != nil {
		log.Fatal("Unable to wirte data")
	}
	// assert content
	Range := ImportConfig().GetGoogleSheetRange()
	if Range == "A2:F" {
		t.Log("TestGetGoogleSheetRange PASS")
	} else {
		t.Error("Content of Range is not correct")
	}
}
