package main

import (
	"TrelloTribbles/internal"
	"log"
	"os"
)

var trelloAPI *internal.Trello

func main() {
	apiKey := os.Getenv("TRELLO_API_KEY")
	if len(apiKey) < 1 {
		log.Fatal("TRELLO_API_KEY is required")
	}

	apiToken := os.Getenv("TRELLO_API_TOKEN")
	if len(apiKey) < 1 {
		log.Fatal("TRELLO_API_TOKEN is required")
	}

	boardId := os.Getenv("TRELLO_BOARD_ID")
	if len(boardId) < 1 {
		log.Fatal("TRELLO_BOARD_ID is required")
	}

	boardSorter := internal.NewBoardSorter(apiKey, apiToken)
	err := boardSorter.SortBoard(boardId)
	if err != nil {
		os.Exit(1)
	}
}
