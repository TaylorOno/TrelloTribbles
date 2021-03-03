package main

import (
	"TrelloTribbles/internal"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/board_sorter", helloHandler)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}