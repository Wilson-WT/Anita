package config

import (
	"os"

	"github.com/go-ini/ini"
	errlog "github.com/inconshreveable/log15"
)

var cfg *ini.File
var _self *Singleton
var configPath = "config.ini"
var eLog = errlog.New("package", "config")

// Singleton ...
type Singleton struct {
}

// CheckAllConfig ...
func CheckAllConfig() {
	ImportConfig().GetGoogleSheetID()
	ImportConfig().GetGoogleSheetRange()
	ImportConfig().GetGoogleSheetTabName()
}

// ImportConfig Import config.ini
func ImportConfig() *Singleton {
	if _self == nil {
		_self = new(Singleton)
	}
	// Load config
	var err error
	cfg, err = ini.InsensitiveLoad(configPath)
	if err != nil {
		eLog.Error("Unable to load " + configPath)
		os.Exit(1)
	}
	return _self
}

// GetGoogleSheetID 取得SheetID
func (o *Singleton) GetGoogleSheetID() string {
	SheetID := cfg.Section("Google Sheets API").Key("SheetID").String()
	if len(SheetID) != 44 {
		eLog.Error("Unable to get SheetID from config.ini")
		os.Exit(1)
	}
	return SheetID
}

// GetGoogleSheetTabName 取得TabName
func (o *Singleton) GetGoogleSheetTabName() string {
	TabName := cfg.Section("Google Sheets API").Key("TabName").String()
	if TabName == "" {
		eLog.Error("Unable to get TabName from config.ini")
		os.Exit(1)
	}
	return TabName
}

// GetGoogleSheetRange 取得Range
func (o *Singleton) GetGoogleSheetRange() string {
	Range := cfg.Section("Google Sheets API").Key("Range").String()
	if Range == "" {
		eLog.Error("Unable to get Range from config.ini")
		os.Exit(1)
	}
	return Range
}
