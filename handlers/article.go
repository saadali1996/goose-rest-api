package handler

import (
	"encoding/json"
	model "github.com/goose-rest-api/models"
	"log"
	"net/http"
)

type test_struct struct {
	Url string `json:"url"`
}

// GetArticle is an httpHandler for route GET /people/{url}
func GetArticle(w http.ResponseWriter, r *http.Request) {

	// swagger:route GET /article article content
	//
	// Get article details against URL
	//
	// This will extract details of an article e.g top_image,
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Params:
	//       url: article_url
	//
	//     Responses:
	//       200: articleResponse
	//       404: jsonError

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)
	decoder := json.NewDecoder(r.Body)

	var t test_struct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	log.Println(t.Url)

	item := model.GetArticleGoose(t.Url)

	if err := json.NewEncoder(w).Encode(item); err != nil {
		panic(err)
	}

}
