package router

import (
	"net/http"

	handler "github.com/goose-rest-api/handlers"
)

// Route type description
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all routes
type Routes []Route

var routes = Routes{
	Route{
		"GetArticle",
		"POST",
		"/article",
		handler.GetArticle,
	},
	Route{
		"GetHealth",
		"GET",
		"/health",
		handler.GetHealth,
	},
}
