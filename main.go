package main

import (
	"log"
	"net/http"
	"os"
	"sort"
)

var trelloAPI = NewTrelloClient(http.DefaultClient, "apiKey", "apiToken")

func main() {
	lists, err := trelloAPI.getBoardLists("boardId")
	if err != nil {
		os.Exit(1)
	}
	sortLists(lists)
}

func sortLists(lists []List) {
	for _, l := range lists {
		Cards, err :=trelloAPI.getListCards(l.Id)
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
		return false
	})
}

func updateCards(cards []Card) {
	for i, c := range cards {
		trelloAPI.updateCardPosition(c.Id, i)
	}
}

