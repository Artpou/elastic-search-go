package controllers

import (
	"log"
	"net/http"

	"github.com/Artpou/elastic-search-go/app/handler/elastic"
	"github.com/Artpou/elastic-search-go/app/handler/respond"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gorilla/mux"
)

type Response struct {
	data string
}

/*func GetElastic() *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	es, err := elasticsearch.NewClient(cfg)

	if err != nil {
		log.Fatalf("Error creating the client : %s", err)
		return nil
	}

	return es
}*/

func GetAllBooks(es *elasticsearch.Client, w http.ResponseWriter, req *http.Request) {
	var books []interface{}
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}
	data := elastic.GetIndex(w, query, "book", es)

	for _, hit := range data.(map[string]interface{})["hits"].([]interface{}) {
		bookData := hit.(map[string]interface{})["_source"]
		books = append(books, bookData)
	}
	respond.RespondJSON(w, 201, books)
}

func DeleteBook(es *elasticsearch.Client, w http.ResponseWriter, r *http.Request) {

}

func UpdateBook(es *elasticsearch.Client, w http.ResponseWriter, r *http.Request) {

}

func CreateBook(es *elasticsearch.Client, w http.ResponseWriter, r *http.Request) {

}

func GetBooksByTitle(es *elasticsearch.Client, w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	title := vars["title"]
	var books []interface{}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": title,
			},
		},
	}

	data := elastic.GetIndex(w, query, "book", es)

	for _, hit := range data.(map[string]interface{})["hits"].([]interface{}) {
		bookData := hit.(map[string]interface{})["_source"]
		//append(books, models.Book{Title: bookData["title"], Author: bookData["author"], Abstract: "ab", Year: string(bookData["year"])})

		books = append(books, bookData)
		/*books = append(books,
			models.Book{
				Title:    bookData(map[string]string)["title"],
				Author:   bookData(map[string]string)["title"],
				Abstract: bookData(map[string]string)["title"],
				Year:     bookData(map[string]string)["title"],
			},
		)*/
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	respond.RespondJSON(w, http.StatusOK, books)
}
