package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"pulley.com/shakesearch/api"

	"pulley.com/shakesearch/model"
)

func startService() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	fmt.Printf("Listening on port %s...", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func setupMrShake() (model.FileLoadSharder, model.Searcher) {
	searcher := model.Searcher{}
	f, err := model.NewFileLoader("completeworks.txt")
	if err != nil {
		log.Fatal(err)
	}
	return f, searcher
}

func main() {
	f, s := setupMrShake()
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/search", api.HandleSearch(f, s))
	startService()
}