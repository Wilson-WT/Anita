package config

import (
	"os"
	"regexp"
	"strconv"
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
	ImportConfig().GetUserName()
	ImportConfig().GetBoardName()
	ImportConfig().GetDoingListName()
	ImportConfig().GetWaitToReportListName()
	ImportConfig().GetReportedListName()
	ImportConfig().GetTrelloKey()
	ImportConfig().GetTrelloToken()
	ImportConfig().GetArchiveReportedCardsWeekDay()
	ImportConfig().GetAutoArchiveTime()
	ImportConfig().GetProjectName()
	ImportConfig().GetSheetServerURL()
	// Check Trello Correspond to Happygorgi Account
	keys := cfg.Section("Email Setting").KeyStrings()
	if len(keys) == 0 {
		eLog.Error("Unable to get any Trello account from config.ini")
		os.Exit(1)
	}
	for _, key := range keys {
		ImportConfig().GetEmailByTrelloUsername(key)
	}
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

// GetUserName 取得UserName
func (o *Singleton) GetUserName() string {
	UserName := cfg.Section("Trello").Key("UserName").String()
	if UserName == "" {
		eLog.Error("Unable to get UserName from config.ini")
		os.Exit(1)
	}
	return UserName
}

// GetBoardName 取得BoardName
func (o *Singleton) GetBoardName() string {
	BoardName := cfg.Section("Trello").Key("BoardName").String()
	if BoardName == "" {
		eLog.Error("Unable to get BoardName from config.ini")
		os.Exit(1)
	}
	return BoardName
}

// GetDoingListName 取得 DoingListName
func (o *Singleton) GetDoingListName() string {
	DoingListName := cfg.Section("Trello").Key("DoingListName").String()
	if DoingListName == "" {
		eLog.Error("Unable to get DoingListName from config.ini")
		os.Exit(1)
	}
	return DoingListName
}

// GetWaitToReportListName 取得WaitToReportListName
func (o *Singleton) GetWaitToReportListName() string {
	WaitToReportListName := cfg.Section("Trello").Key("WaitToReportListName").String()
	if WaitToReportListName == "" {
		eLog.Error("Unable to get WaitToReportListName from config.ini")
		os.Exit(1)
	}
	return WaitToReportListName
}

// GetReportedListName 取得ReportedListName
func (o *Singleton) GetReportedListName() string {
	ReportedListName := cfg.Section("Trello").Key("ReportedListName").String()
	if ReportedListName == "" {
		eLog.Error("Unable to get ReportedListName from config.ini")
		os.Exit(1)
	}
	return ReportedListName
}

// GetTrelloKey 取得Key
func (o *Singleton) GetTrelloKey() string {
	Key := cfg.Section("Trello").Key("Key").String()
	if len(Key) != 32 {
		eLog.Error("Unable to get Key from config.ini")
		os.Exit(1)
	}
	return Key
}

// GetTrelloToken 取得Token
func (o *Singleton) GetTrelloToken() string {
	Token := cfg.Section("Trello").Key("Token").String()
	if len(Token) != 64 {
		eLog.Error("Unable to get Token from config.ini")
		os.Exit(1)
	}
	return Token
}

// GetArchiveReportedCardsWeekDay 取得ArchiveReportedCardsWeekDay
func (o *Singleton) GetArchiveReportedCardsWeekDay() string {
	ArchiveReportedCardsWeekDay := cfg.Section("Trello").Key("ArchiveReportedCardsWeekDay").String()
	if ArchiveReportedCardsWeekDay == "" {
		return ""
	}
	splitedWeekDayStrings := strings.Split(ArchiveReportedCardsWeekDay, ",")
	for _, weekDayString := range splitedWeekDayStrings {
		weekDay, err := strconv.ParseInt(weekDayString, 10, 64)
		if err != nil || weekDay < 0 || weekDay > 6 {
			eLog.Error("Unable to get ArchiveReportedCardsWeekDay from config.ini")
			os.Exit(1)
		}
	}
	return ArchiveReportedCardsWeekDay
}

// GetAutoArchiveTime 取得AutoArchiveTime
func (o *Singleton) GetAutoArchiveTime() string {
	AutoArchiveTime := cfg.Section("Trello").Key("AutoArchiveTime").String()
	// 0:00 or 00:00 to 23:59
	timeChecker := regexp.MustCompile(`^([0-9]|0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$`)
	if timeChecker.MatchString(AutoArchiveTime) {
		return AutoArchiveTime
	}
	eLog.Error("Unable to get AutoArchiveTime from config.ini")
	os.Exit(1)
	return ""
}

// GetEmailByTrelloUsername 取得Email
func (o *Singleton) GetEmailByTrelloUsername(trelloUsername string) string {
	Account := cfg.Section("Email Setting").Key(trelloUsername).String()
	if len(Account) == 0 {
		eLog.Error("Unable to get LDAP account from config.ini")
		os.Exit(1)
	}
	return Account + "@happygorgi.com"
}

// GetProjectName 取得Team
func (o *Singleton) GetProjectName() string {
	projectNameInConfig := cfg.Section("Project Setting").Key("ProjectName").String()
	projectNameArray := []string{"AEP", "AFU", "HBC", "TF", "Research", "IT", "Sales/Presales", "Marketing/PR", "HR (含 interview)", "Admin"}
	for _, projectName := range projectNameArray {
		if projectName == projectNameInConfig {
			return projectNameInConfig
		}
	}
	eLog.Error("Unable to get project name from config.ini")
	os.Exit(1)
	return ""
}

// GetSheetServerURL 取得 SheetServerURL
func (o *Singleton) GetSheetServerURL() string {
	return "http://10.20.108.20:5000"
}
