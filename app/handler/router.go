package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	log.Println("Starting development server at http://127.0.0.1:10000/")
	router := mux.NewRouter().StrictSlash(true)
	es, err := elasticsearch.NewDefaultClient()

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	log.Println(elasticsearch.Version)

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	log.Println(res)

	router.HandleFunc("/", HomeRoute)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":10000", router))
}

func HomeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to our API")
}
