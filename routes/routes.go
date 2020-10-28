package routes

import (
	"net/http"

	handler "github.com/TMDB/handlers"
)

//Init ...
func Init() {
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/search", handler.Dashboard)
	http.HandleFunc("/search/person", handler.RoutePerson)
	http.HandleFunc("/search/person/knownfor", handler.RouteKnownFor)
	http.HandleFunc("/search/movie", handler.RouteMovie)
	http.HandleFunc("/search/movie/popular", handler.RoutePopularMovie)
	http.HandleFunc("/search/tvshow", handler.RouteTVShow)
	http.HandleFunc("/search/tvshow/popular", handler.RoutePopularTV)
	http.HandleFunc("/search/company", handler.RouteCompany)
	http.HandleFunc("/search/keyword", handler.RouteKeyWord)

	http.Handle("/favicon.ico", http.NotFoundHandler())
}
