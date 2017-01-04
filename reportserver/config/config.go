package config

import (
	"os"
	"strings"

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
	ImportConfig().GetSheetServerURL()
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

// GetSheetServerURL 取得SheetServerURL
func (o *Singleton) GetSheetServerURL() string {
	SheetServerURL := cfg.Section("Sheet Server").Key("URL").String()
	if strings.HasPrefix(SheetServerURL, "http://") {
		return SheetServerURL
	}
	eLog.Error("Unable to get SheetServerUrl from config.ini")
	os.Exit(1)
	return ""
}
