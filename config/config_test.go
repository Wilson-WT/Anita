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
	d1 := []byte("[Trello]\n" +
		"UserName = username\n" +
		"BoardName = AFU QA Team - Kanban Board\n" +
		"DoingListName = Doing\n" +
		"WaitToReportListName = Wait To Report To Form\n" +
		"ReportedListName = Reported\n" +
		"Key = test8cff11fcd85e1cb6b5822840a123\n" +
		"Token = test2729b6e332d6c31222da5053601ab4b3ae314fe209db2c3b768ac538bd39\n" +
		"ArchiveReportedCardsWeekDay = 1,3,5\n" +
		"AutoArchiveTime = 18:00\n\n" +
		"[Email Setting]\n" +
		"trello1 = account1\n" +
		"trello2 = account2\n" +
		"trello3 = account3\n\n" +
		"[Project Setting]\n" +
		"ProjectName = AFU\n\n")
	err := ioutil.WriteFile(configPath, d1, 0644)
	if err != nil {
		log.Fatal("Unable to wirte data")
	}

	// check all config
	CheckAllConfig()
}

//// TestGetGetUserName
func TestGetUserName_Valid(t *testing.T) {
	configPath = "config_test.ini"

	// Prepare test data
	d1 := []byte("[Trello]\nUserName = qatest")
	err := ioutil.WriteFile(configPath, d1, 0644)
	if err != nil {
		log.Fatal("Unable to wirte data")
	}

	// assert content
	userName := ImportConfig().GetUserName()
	if userName == "qatest" {
		t.Log("TestGetUserName PASS")
	} else {
		t.Error("Content of UserName is not correct")
	}
}

//// TestGetBoardName
func TestGetBoardName_Valid(t *testing.T) {
	configPath = "config_test.ini"

	// Prepare test data
	d1 := []byte("[Trello]\nBoardName = qatest")
	err := ioutil.WriteFile(configPath, d1, 0644)
	if err != nil {
		log.Fatal("Unable to wirte data")
	}

	// assert content
	BoardName := ImportConfig().GetBoardName()
	if BoardName == "qatest" {
		t.Log("TestGetBoardName PASS")
	} else {
		t.Error("Content of BoardName is not correct")
	}
}

//// TestGetDoingListName
func TestGetDoingListName_Valid(t *testing.T) {
	configPath = "config_test.ini"

	// Prepare test data
	d1 := []byte("[Trello]\nDoingListName = Doing")
	err := ioutil.WriteFile(configPath, d1, 0644)
	if err != nil {
		log.Fatal("Unable to wirte data")
	}

	// assert content
	DoingListName := ImportConfig().GetDoingListName()
	if DoingListName == "Doing" {
		t.Log("TestGetDoingListName PASS")
	} else {
		t.Error("Content of DoingListName is not correct")
	}
}

//// TestGetWaitToReportListName
func TestGetWaitToReportListName_Valid(t *testing.T) {
	configPath = "config_test.ini"

	// Prepare test data
	d1 := []byte("[Trello]\nWaitToReportListName = qatest")
	err := ioutil.WriteFile(configPath, d1, 0644)
	if err != nil {
		log.Fatal("Unable to wirte data")
	}

	// assert content
	WaitToReportListName := ImportConfig().GetWaitToReportListName()
	if WaitToReportListName == "qatest" {
		t.Log("TestGetWaitToReportListName PASS")
	} else {
		t.Error("Content of WaitToReportListName is not correct")
	}
}

//// TestGetReportedListName
func TestGetReportedListName_Valid(t *testing.T) {
	configPath = "config_test.ini"

	// Prepare test data
	d1 := []byte("[Trello]\nReportedListName = qatest")
	err := ioutil.WriteFile(configPath, d1, 0644)
	if err != nil {
		log.Fatal("Unable to wirte data")
	}

	// assert content
	ReportedListName := ImportConfig().GetReportedListName()
	if ReportedListName == "qatest" {
		t.Log("TestGetReportedListName PASS")
	} else {
		t.Error("Content of ReportedListName is not correct")
	}
}

//// TestGetReportedListName
func TestGetTrelloKey_Valid(t *testing.T) {
	configPath = "config_test.ini"

	// Prepare test data
	d1 := []byte("[Trello]\nKey = 3b258cff11fcd85e1cb6b5822840aca7")
	err := ioutil.WriteFile(configPath, d1, 0644)
	if err != nil {
		log.Fatal("Unable to wirte data")
	}

	// assert content
	TrelloKey := ImportConfig().GetTrelloKey()
	if TrelloKey == "3b258cff11fcd85e1cb6b5822840aca7" {
		t.Log("TestGetTrelloKey PASS")
	} else {
		t.Error("Content of TrelloKey is not correct")
	}
}

//// TestGetTrelloToken
func TestGetTrelloToken_Valid(t *testing.T) {
	configPath = "config_test.ini"

	// Prepare test data
	d1 := []byte("[Trello]\nToken = ff7a2729b6e332d6c31222da5053601ab4b3ae314fe209db2c3b768ac538bxxx")
	err := ioutil.WriteFile(configPath, d1, 0644)
	if err != nil {
		log.Fatal("Unable to wirte data")
	}

	// assert content
	TrelloToken := ImportConfig().GetTrelloToken()
	if TrelloToken == "ff7a2729b6e332d6c31222da5053601ab4b3ae314fe209db2c3b768ac538bxxx" {
		t.Log("TestGetTrelloToken PASS")
	} else {
		t.Error("Content of TrelloToken is not correct")
	}
}

//// TestGetEmailByTrelloUsername
func TestGetEmailByTrelloUsername_Valid(t *testing.T) {
	configPath = "config_test.ini"

	// Prepare test data
	d1 := []byte("[Email Setting]\ntrellouser = happygorgi.account")
	err := ioutil.WriteFile(configPath, d1, 0644)
	if err != nil {
		log.Fatal("Unable to wirte data")
	}

	// assert content
	Email := ImportConfig().GetEmailByTrelloUsername("trellouser")
	if Email == "happygorgi.account@happygorgi.com" {
		t.Log("TestGetEmailByTrelloUsername PASS")
	} else {
		t.Error("Content of trello username or account is not correct")
	}
}

//// TestGetArchiveReportedCardsWeekDay
func TestGetArchiveReportedCardsWeekDay_Valid(t *testing.T) {
	configPath = "config_test.ini"

	testCases := []struct {
		archiveReportedCardsWeekDay string
	}{
		{""},
		{"0"},
		{"0,1"},
		{"0,1,2"},
		{"0,1,2,3"},
		{"0,1,2,3,4"},
		{"0,1,2,3,4,5"},
		{"0,1,2,3,4,5,6"},
		{"0,1,3,4,5,6"},
		{"0,1,3,5,6"},
		{"0,1,3,6"},
		{"0,1,3"},
		{"0,3"},
		{"3"},
	}

	for _, testCase := range testCases {
		testArchiveReportedCardsWeekDay := testCase.archiveReportedCardsWeekDay
		// Prepare test data
		d1 := []byte("[Trello]\nArchiveReportedCardsWeekDay = " + testArchiveReportedCardsWeekDay)
		err := ioutil.WriteFile(configPath, d1, 0644)
		if err != nil {
			log.Fatal("Unable to wirte data")
		}

		// assert content
		ArchiveReportedCardsWeekDay := ImportConfig().GetArchiveReportedCardsWeekDay()
		if ArchiveReportedCardsWeekDay == testArchiveReportedCardsWeekDay {
			t.Log("TestArchiveReportedCardsWeekDay PASS")
		} else {
			t.Error("Content of ArchiveReportedCardsWeekDay is not correct")
		}
	}
}

//// TestGetAutoArchiveTime
func TestGetAutoArchiveTime_Valid(t *testing.T) {
	configPath = "config_test.ini"

	testCases := []struct {
		autoArchiveTime string
	}{
		{"0:00"},
		{"00:00"},
		{"01:00"},
		{"1:00"},
		{"23:59"},
	}
	for _, testCase := range testCases {
		testAutoArchiveTime := testCase.autoArchiveTime
		// Prepare test data
		d1 := []byte("[Trello]\nAutoArchiveTime = " + testAutoArchiveTime)
		err := ioutil.WriteFile(configPath, d1, 0644)
		if err != nil {
			log.Fatal("Unable to wirte data")
		}

		// assert content
		actualAutoArchiveTime := ImportConfig().GetAutoArchiveTime()
		if actualAutoArchiveTime == testAutoArchiveTime {
			t.Log("TestGetAutoArchiveTime PASS")
		} else {
			t.Error("Content of AutoArchiveTime is not correct")
		}
	}
}

//// TestGetProjectName
func TestGetProjectName_Valid(t *testing.T) {
	configPath = "config_test.ini"

	testCases := []struct {
		projectName string
	}{
		{"AEP"},
		{"AFU"},
		{"HBC"},
		{"TF"},
		{"Research"},
		{"IT"},
		{"Sales/Presales"},
		{"Marketing/PR"},
		{"HR (含 interview)"},
		{"Admin"},
	}
	for _, testCase := range testCases {
		// Prepare test data
		testProjectName := testCase.projectName
		d1 := []byte("[Project Setting]\nProjectName = " + testProjectName)
		err := ioutil.WriteFile(configPath, d1, 0644)
		if err != nil {
			log.Fatal("Unable to wirte data")
		}
		// assert content
		ProjectName := ImportConfig().GetProjectName()
		if ProjectName == testProjectName {
			t.Log("TestGetProjectName PASS")
		} else {
			t.Error("Content of ProjectName is not correct")
		}
	}
}

//// TestGetSheetServerURL
func TestGetSheetServerURL_Valid(t *testing.T) {
	// assert content
	SheetServerURL := ImportConfig().GetSheetServerURL()
	if SheetServerURL == "http://10.20.108.20:5000" {
		t.Log("TestGetSheetServerURL PASS")
	} else {
		t.Error("Content of SheetServerURL is not correct")
	}
}
