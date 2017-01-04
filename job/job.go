package job

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ibmboy19/Anita/config"
	"github.com/ibmboy19/Anita/trello"
	trelloApi "github.com/ibmboy19/go-trello"
	errlog "github.com/inconshreveable/log15"
)

var eLog = errlog.New("package", "job")

func getCardWorkingHours(card trelloApi.Card) string {
	var filteredActions []trelloApi.Action
	actions, err := card.Actions()
	if err != nil {
		eLog.Error("Get card Action error!")
	}
	for _, action := range actions {
		if action.Data.ListBefore.Name != "" && action.Data.ListAfter.Name != "" {
			filteredActions = append(filteredActions, action)
		} else if action.Type == "createCard" && action.Data.List.Name == config.ImportConfig().GetDoingListName() {
			filteredActions = append(filteredActions, action)
		}
	}

	// Calculate working hours for card
	var workingHours float64
	for index, action := range filteredActions {
		if action.Data.ListAfter.Name == config.ImportConfig().GetDoingListName() || action.Type == "createCard" {
			nextAction := filteredActions[index-1]
			nextActionTime, err1 := time.Parse(time.RFC3339, nextAction.Date)
			actionTime, err2 := time.Parse(time.RFC3339, action.Date)
			if err1 != nil || err2 != nil {
				eLog.Error("Parse card action time error!")
				continue
			} else {
				// convert to local time
				nextActionTime = nextActionTime.Local()
				actionTime = actionTime.Local()
			}
			workingHoursString := calculateWorkingHour(actionTime, nextActionTime)
			tmpWorkingHour, _ := strconv.ParseFloat(workingHoursString, 64)
			workingHours += tmpWorkingHour
		}
	}
	return strconv.FormatFloat(workingHours, 'f', 1, 64)
}

func getCardDoneDate(card trelloApi.Card) string {
	// get card at Doing's timestamp and at Done's timestamp
	actions, err := card.Actions()
	if err != nil {
		eLog.Error("Get card Action error!")
	}
	var doneAction trelloApi.Action
	var doneTime time.Time
	for _, action := range actions {
		listAfterName := action.Data.ListAfter.Name
		if listAfterName == config.ImportConfig().GetWaitToReportListName() {
			doneAction = action
			break
		}
	}

	// Doing Time and Done time : 2016-11-09T09:09:22.900Z
	doneTime, err1 := time.Parse(time.RFC3339, doneAction.Date)
	if err1 != nil {
		eLog.Error("Parse card action time error!")
	} else {
		// convert to local time
		doneTime = doneTime.Local()
	}
	return doneTime.Format("2006/1/2")
}

func getFormattedCardName(cardName string) string {
	// replace " with \" to fit json string format
	return strings.Replace(cardName, "\"", "\\\"", -1)
}

// DoTrelloReportJob 執行 trello 工作週報任務
func DoTrelloReportJob() {
	// Get cards from trello
	cards := trello.GetCardsInBoardList(config.ImportConfig().GetUserName(), config.ImportConfig().GetBoardName(), config.ImportConfig().GetWaitToReportListName())
	for _, card := range cards {
		log.Println(" ----------------------- ")
		log.Println("Card Name: ", card.Name)
		// Record if this card reported successfully
		var isReportedSuccess = true
		// Store card done date
		var cardDoneDate string

		// Get working hours string
		workingHoursString := getCardWorkingHours(card)
		log.Println("Card hour: ", workingHoursString)

		// Get card done date
		cardDoneDate = getCardDoneDate(card)
		log.Println("Card done date: " + cardDoneDate)

		// Get members of the card
		members, _ := card.Members()
		for _, member := range members {
			log.Println("Card members: ", member.FullName)
			// Getting User Email
			userEmail := config.ImportConfig().GetEmailByTrelloUsername(member.Username)

			// Time formating - 2016/7/4 下午 7:21:07
			currentTimeStamp := time.Now().Format("2006/1/2 pm 3:04:05")
			currentTimeStamp = strings.Replace(currentTimeStamp, "am", "上午", 1)
			currentTimeStamp = strings.Replace(currentTimeStamp, "pm", "下午", 1)

			// Getting project name
			project := config.ImportConfig().GetProjectName()

			formattedCardName := getFormattedCardName(card.Name)
			// Call sheet server's API to Append Sheet
			jsonString := `{"dateTime":"` + currentTimeStamp + `", "user":"` + userEmail + `", "project":"` + project + `",
		                "content":"` + formattedCardName + `", "hours":"` + workingHoursString + `", "finishDate":"` + cardDoneDate + `"}`
			response, err := doSheetServerPOST(jsonString)

			if err != nil {
				isReportedSuccess = false
				eLog.Error("Unable to reach the server.")
			} else {
				// Error handling
				if response.StatusCode == http.StatusOK {
					log.Println("Report Card successfully")
				} else {
					isReportedSuccess = false
					eLog.Error("Report Card failed")
				}
			}
		}
		// Drag card to Reported List
		if isReportedSuccess {
			listID := trello.GetListID(config.ImportConfig().GetUserName(), config.ImportConfig().GetBoardName(), config.ImportConfig().GetReportedListName())
			card.MoveToList(listID)
			card.AddComment("By Working report automation : REPORTED!!!!!")
		}
	}
}

func doSheetServerPOST(json string) (*http.Response, error) {
	sheetServerURL := config.ImportConfig().GetSheetServerURL()

	req, _ := http.NewRequest("POST", sheetServerURL, bytes.NewBuffer([]byte(json)))
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := http.Client{}
	response, err := client.Do(req)

	return response, err
}

// ArchiveReportedCards () no return
func ArchiveReportedCards() {
	// Get cards at "Reported" phase
	cards := trello.GetCardsInBoardList(config.ImportConfig().GetUserName(), config.ImportConfig().GetBoardName(), config.ImportConfig().GetReportedListName())
	for _, card := range cards {
		card.Archive()
	}
}
