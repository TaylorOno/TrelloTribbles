//go:generate mockgen -destination=mocks/mock_trello.go -package=mocks -source trello.go

package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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
	Id        string   `json:id`
	Name      string   `json:name`
	Pos       float64  `json:pos`
	IdMembers []string `json:idMembers`
}

func NewTrelloClient(client Client, apiKey string, apiToken string) *Trello {
	return &Trello{
		client:   client,
		apiKey:   apiKey,
		apiToken: apiToken,
	}
}

func (t *Trello) getBoardLists(boardId string) ([]List, error) {
	var lists []List

	url := fmt.Sprintf("https://api.trello.com/1/boards/%v/lists", boardId)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	err := t.trelloRequest(req, &lists)
	if err != nil {
		return lists, err
	}

	return lists, err
}

func (t *Trello) getListCards(listId string) ([]Card, error) {
	var cards []Card

	url := fmt.Sprintf("https://api.trello.com/1/lists/%v/cards", listId)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	err := t.trelloRequest(req, &cards)
	if err != nil {
		return cards, err
	}

	return cards, err
}

func (t *Trello) updateCardPosition(cardId string, position int) error {
	var card Card

	url := fmt.Sprintf("https://api.trello.com/1/cards/%v", cardId)
	reqBody := fmt.Sprintf(`{"pos":%v}`, position)
	req, _ := http.NewRequest(http.MethodPut, url, strings.NewReader(reqBody))

	return t.trelloRequest(req, &card)
}

func (t *Trello) trelloRequest(req *http.Request, cards interface{}) error {
	t.setAuth(req)
	resp, err := t.client.Do(req)
	if err != nil {
		return err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &cards)
}

func (t *Trello) setAuth(req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("OAuth oauth_consumer_key=\"%v\", oauth_token=\"%v\"", t.apiKey, t.apiToken))
	req.Header.Set("Content-Type", "application/json")
}
