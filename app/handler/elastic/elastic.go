package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Artpou/elastic-search-go/app/handler/respond"
	"github.com/Artpou/elastic-search-go/app/views"
	"github.com/elastic/go-elasticsearch/v7"
)

var es *elasticsearch.Client
var r map[string]interface{}

func GetElasticSearchClient() *elasticsearch.Client {
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
}

func GetIndex(w http.ResponseWriter, query map[string]interface{}, index string, es *elasticsearch.Client) interface{} {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(index),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			respond.RespondJSON(w, http.StatusNotFound, views.BodyParsingError())
		} else {
			// Print the response status and error information.
			respond.RespondJSON(w, http.StatusNotFound, e["error"].(map[string]interface{})["reason"])
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	/*// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)*/
	// Print the ID and document source for each hit.

	return r["hits"]
}
