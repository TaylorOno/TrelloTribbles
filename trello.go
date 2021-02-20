//go:generate mockgen -destination=mocks/mock_trello.go -package=mocks -source trello.go

package main

import (
	"net/http"
)

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

type Trello struct {
	client   Client
	apiKey   string
	apiToken string
}

type List struct {
	Id string `json:id`
}

type Card struct {
	Id string `json:id`
}

func NewTrelloClient(client Client, apiKey string, apiToken string) *Trello {
	return &Trello{
		client:   client,
		apiKey:   apiKey,
		apiToken: apiToken,
	}
}

func (t *Trello) getBoardLists(boardId string) ([]List, error) {
	lists := make([]List, 0)
	return lists, nil
}

func (t *Trello) getListCards(listId string) ([]Card, error) {
	cards := make([]Card, 0)
	return cards, nil
}

func (t *Trello) updateCardPosition(cardId string, position int) error {
	return nil
}
