package router

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/Artpou/elastic-search-go/app/controllers"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gorilla/mux"
)

var (
	r  map[string]interface{}
	wg sync.WaitGroup
	es *elasticsearch.Client
)

func HandleRequests() {
	log.Println(strings.Repeat("~", 37))
	log.Println("Starting development server at http://127.0.0.1:8080/")

	router := mux.NewRouter().StrictSlash(true)
	//db = database.InitDb()

	router.HandleFunc("/", HomeLink)
	router.HandleFunc("/api/book/", GetBooks).Methods("GET")

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":8080", router))
}

//Book
func GetBooks(w http.ResponseWriter, req *http.Request) {
	controllers.GetAllBooks(w, req)
}

//LOGIN
func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}
