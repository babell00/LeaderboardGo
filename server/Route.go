package server

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Pattern     string
	Method      string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"ShowLeaderboard", "/leaderboard", "GET", ShowLeaderboard},
	Route{"AddPlayerScore", "/playerscore", "POST", AddPlayerScore},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).HandlerFunc(route.HandlerFunc)
	}
	return router
}
