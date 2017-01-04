package config

import "testing"

func TestGetCronSettingForArchiveReportedCards_Valid(t *testing.T) {
	// week day = "1,2,3", daily hour = "9"
	archiveReportedCardsCronSetting := GetCronSettingForArchiveReportedCards("1,2,3", "9:00")
	if archiveReportedCardsCronSetting == "0 0 9 * * 1,2,3" {
		t.Log("GetCronSettingForArchiveReportedCards PASS")
	} else {
		t.Error("GetCronSettingForArchiveReportedCards FAIL")
	}

	// week day = "1,2,3", daily hour = "24"
	archiveReportedCardsCronSetting = GetCronSettingForArchiveReportedCards("1,2,3", "24:00")
	if archiveReportedCardsCronSetting == "0 0 0 * * 1,2,3" {
		t.Log("GetCronSettingForArchiveReportedCards PASS")
	} else {
		t.Error("GetCronSettingForArchiveReportedCards FAIL")
	}
}
