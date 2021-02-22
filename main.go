package main

import (
	"log"
	"net/http"
	"os"
	"sort"
)

var trelloAPI *Trello

func main() {
	apiKey := os.Getenv("TRELLO_API_KEY")
	if len(apiKey) < 1 {
		log.Fatal("TRELLO_API_KEY is required")
	}

	apiToken := os.Getenv("TRELLO_API_TOKEN")
	if len(apiKey) < 1 {
		log.Fatal("TRELLO_API_TOKEN is required")
	}

	trelloAPI = NewTrelloClient(http.DefaultClient, apiKey, apiToken)

	boardId := os.Getenv("TRELLO_BOARD_ID")
	lists, err := trelloAPI.getBoardLists(boardId)
	if err != nil {
		os.Exit(1)
	}
	sortLists(lists)
}

func sortLists(lists []List) {
	for _, l := range lists {
		Cards, err := trelloAPI.getListCards(l.Id)
		if err != nil {
			log.Printf("failed to get cards for list: %v", l.Id)
			break
		}
		sortCards(Cards)
		updateCards(Cards)
	}
}

func sortCards(cards []Card) {
	sort.Slice(cards, func(i, j int) bool {
		if len(cards[i].IdMembers) > len(cards[j].IdMembers) {
			return true
		}
		return false
	})
}

func updateCards(cards []Card) {
	for i, c := range cards {
		err := trelloAPI.updateCardPosition(c.Id, 1<<i)
		if err != nil {
			log.Printf("failed to update card: %v\n", err)
			return
		}
	}
}
