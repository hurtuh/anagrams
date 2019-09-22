package routes

import (
	"github.com/anagrams/service"
	"github.com/gorilla/mux"
)

type Routers struct {
	Handlers *service.Service
}

func (rout *Routers) MakeHTTPHandlers() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("POST").Path("/set_anagrams").HandlerFunc(rout.Handlers.SetAnagrams)
	router.Methods("GET").Path("/search_anagrams").HandlerFunc(rout.Handlers.SearchAnagrams)

	return router
}