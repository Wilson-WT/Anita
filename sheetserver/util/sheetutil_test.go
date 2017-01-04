package util

import "testing"

func TestGetWorkingReports_Invalid(t *testing.T) {
	// accountName = "", period = "201608-201610"
	reports, err := GetWorkingReports("", "201608-201610")
	if reports == nil && err.Error() == "Account name is invalid" {
		t.Log("GetWorkingReports PASS")
	} else {
		t.Error("GetWorkingReports FAIL")
	}
	// accountName = "hello.world", period = ""
	reports, err = GetWorkingReports("hello.world", "")
	if reports == nil && err.Error() == "Period is invalid" {
		t.Log("GetWorkingReports PASS")
	} else {
		t.Error("GetWorkingReports FAIL")
	}
	// accountName = "hello.world", period = "201608-201607"
	reports, err = GetWorkingReports("hello.world", "201608-201607")
	if reports == nil && err.Error() == "Period is invalid" {
		t.Log("GetWorkingReports PASS")
	} else {
		t.Error("GetWorkingReports FAIL")
	}
	// accountName = "hello.world", period = "201708-201609"
	reports, err = GetWorkingReports("hello.world", "201708-201609")
	if reports == nil && err.Error() == "Period is invalid" {
		t.Log("GetWorkingReports PASS")
	} else {
		t.Error("GetWorkingReports FAIL")
	}
	// accountName = "hello.world", period = "201613-201614"
	reports, err = GetWorkingReports("hello.world", "201613-201614")
	if reports == nil && err.Error() == "Period is invalid" {
		t.Log("GetWorkingReports PASS")
	} else {
		t.Error("GetWorkingReports FAIL")
	}
	// accountName = "hello.world", period = "20160-201612"
	reports, err = GetWorkingReports("hello.world", "2016ab-201612")
	if reports == nil && err.Error() == "Period is invalid" {
		t.Log("GetWorkingReports PASS")
	} else {
		t.Error("GetWorkingReports FAIL")
	}
}
