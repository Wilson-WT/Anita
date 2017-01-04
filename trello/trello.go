package trello

import (
	"os"

	"github.com/ibmboy19/Anita/config"
	"github.com/ibmboy19/go-trello"
	errlog "github.com/inconshreveable/log15"
)

var eLog = errlog.New("package", "trello")

// GetCardsInBoardList (userName, boeardName, listName string) []trello.Card
func GetCardsInBoardList(userName, boeardName, listName string) []trello.Card {
	user := getUser(userName)
	board := getBoard(user, boeardName)
	list := getList(board, listName)
	cards := getCards(list)
	return cards
}

// GetListID 取得 List ID
func GetListID(userName, boardName, listName string) string {
	user := getUser(userName)
	board := getBoard(user, boardName)
	list := getList(board, listName)
	return list.Id
}

func getUser(userName string) *trello.Member {
	trelloToken := config.ImportConfig().GetTrelloToken()
	// Client
	trelloClient, err := trello.NewAuthClient(config.ImportConfig().GetTrelloKey(), &trelloToken)
	if err != nil {
		eLog.Error(err.Error())
		os.Exit(1)
	}
	user, err := trelloClient.Member(userName)
	if err != nil {
		eLog.Error(err.Error())
		os.Exit(1)
	}
	return user
}

func getBoard(user *trello.Member, boeardName string) *trello.Board {
	boards, err := user.Boards()
	if err != nil {
		eLog.Error(err.Error())
		os.Exit(1)
	}
	// Get Specific board id
	var wantedBoard trello.Board
	for _, board := range boards {
		if board.Name == boeardName {
			wantedBoard = board
		}
	}
	return &wantedBoard
}

func getList(board *trello.Board, listName string) *trello.List {
	lists, err := board.Lists()
	if err != nil {
		eLog.Error(err.Error())
		os.Exit(1)
	}
	// Get specific list
	var wantList trello.List
	for _, list := range lists {
		if list.Name == listName {
			wantList = list
		}
	}
	return &wantList
}

func getCards(list *trello.List) []trello.Card {
	cards, _ := list.Cards()
	return cards
}
