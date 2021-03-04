package internal

import (
	"log"
	"net/http"
	"sort"
)

type BoardSorter struct {
	Trello *Trello
}

//NewBoardSorter give an apiKey and apiToken String returns a new board sorter
func NewBoardSorter(apiKey string, apiToken string) *BoardSorter {
	return &BoardSorter{
		Trello: NewTrelloClient(http.DefaultClient, apiKey, apiToken),
	}
}

//SortBoard sorts all lists on the given board id
func (b *BoardSorter) SortBoard(id string) error {
	lists, err := b.Trello.getBoardLists(id)
	if err != nil {
		return err
	}

	b.sortLists(lists)
	return nil
}

//sortLists will loop through a given
func (b *BoardSorter) sortLists(lists []List) {
	for _, l := range lists {
		b.sortList(l)
	}
}

//sortList will sort all cards on a given list and update the card positions
func (b *BoardSorter) sortList(l List) {
	Cards, err := b.Trello.getListCards(l.Id)
	if err != nil {
		log.Printf("failed to get cards for list: %v", l.Id)
		return
	}
	b.sortCards(Cards)
	b.updateCards(Cards)
}

//sortCards will order an array of cards with in descending order based on members following
func (b *BoardSorter) sortCards(cards []Card) {
	sort.Slice(cards, func(i, j int) bool {
		if len(cards[i].IdMembers) > len(cards[j].IdMembers) {
			return true
		}
		return false
	})
}

//updateCards updates a cards position to match the order in the sorted list
func (b *BoardSorter) updateCards(cards []Card) {
	for i, c := range cards {
		err := b.Trello.updateCardPosition(c.Id, 1<<i)
		if err != nil {
			log.Printf("failed to update card: %v\n", err)
			return
		}
	}
}
