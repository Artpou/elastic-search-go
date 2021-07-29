package router

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Artpou/elastic-search-go/app/controllers"
	"github.com/Artpou/elastic-search-go/app/handler/elastic"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gorilla/mux"
)

var es *elasticsearch.Client

func HandleRequests() {
	log.Println(strings.Repeat("~", 37))
	log.Println("Starting development server at http://127.0.0.1:8080/")

	es = elastic.GetElasticSearchClient()
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", HomeLink)
	router.HandleFunc("/api/book/", GetBooks).Methods("GET")
	router.HandleFunc("/api/book/{title}", GetBooksByTitle).Methods("GET")

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":8080", router))
}

// Book
func GetBooks(w http.ResponseWriter, req *http.Request) {
	controllers.GetAllBooks(es, w, req)
}

func GetBooksByTitle(w http.ResponseWriter, req *http.Request) {
	controllers.GetBooksByTitle(es, w, req)
}

func AddBook(w http.ResponseWriter, req *http.Request) {
	controllers.CreateBook(es, w, req)
}

func RemoveBook(w http.ResponseWriter, req *http.Request) {
	controllers.DeleteBook(es, w, req)
}

func UpdateBook(w http.ResponseWriter, req *http.Request) {
	controllers.UpdateBook(es, w, req)
}

// Home
func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}
