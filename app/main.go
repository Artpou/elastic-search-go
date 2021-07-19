package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/elastic/go-elasticsearch/v7"
)

func main() {
	fmt.Println("Hello")
	es, _ := elasticsearch.NewDefaultClient()
	log.Println(elasticsearch.Version)
	log.Println(es.Info())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
