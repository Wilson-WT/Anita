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
	d1 := []byte("[Sheet Server]\n" +
		"URL = http://10.20.108.20:5000")
	err := ioutil.WriteFile(configPath, d1, 0644)
	if err != nil {
		log.Fatal("Unable to wirte data")
	}

	// check all config
	CheckAllConfig()
}

//// TestGetSheetServerURL
func TestGetSheetServerURL_Valid(t *testing.T) {
	configPath = "config_test.ini"

	// Prepare test data
	d1 := []byte("[Sheet Server]\nURL = http://10.20.108.20:5000")
	err := ioutil.WriteFile(configPath, d1, 0644)
	if err != nil {
		log.Fatal("Unable to wirte data")
	}

	// assert content
	SheetServerURL := ImportConfig().GetSheetServerURL()
	if SheetServerURL == "http://10.20.108.20:5000" {
		t.Log("TestGetSheetServerURL PASS")
	} else {
		t.Error("Content of SheetServerURL is not correct")
	}
}
