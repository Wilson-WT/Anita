package main

import (
	"github.com/ibmboy19/Anita/config"
	"github.com/ibmboy19/Anita/job"
	"github.com/robfig/cron"
)

/**
 * Working Report Automation
 * Author: Arkflex QA Team
 * Finish date: 2016/08/24
 */
func main() {
	// Check all config settings are correct
	config.CheckAllConfig()

	// Cron job
	cron := cron.New()
	// Get cron setting for archiving reported cards
	archiveReportedCardsWeekDay := config.ImportConfig().GetArchiveReportedCardsWeekDay()
	// Add cron jobs
	cron.AddFunc("0 30 23 * * *", job.DoTrelloReportJob) // DoTrelloReportJob at 23:30 everyday.
	if archiveReportedCardsWeekDay != "" {
		autoArchiveTime := config.ImportConfig().GetAutoArchiveTime()
		archiveReportedCardsCronSetting := config.GetCronSettingForArchiveReportedCards(archiveReportedCardsWeekDay, autoArchiveTime)
		cron.AddFunc(archiveReportedCardsCronSetting, job.ArchiveReportedCards) // ArchiveReportedCards after daily standup 1 hour (18:00) every MON,WED and FRI.
	}
	cron.Start()
	// pause the process
	select {}
}
