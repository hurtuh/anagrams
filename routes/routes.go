package routes

import (
	"github.com/hurtuh/anagrams/service"
	"github.com/gorilla/mux"
)

type Routers struct {
	Handlers *service.Service
}

func (rout *Routers) MakeHTTPHandlers() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("POST").Path("/load").HandlerFunc(rout.Handlers.SetAnagrams)
	router.Methods("GET").Path("/get").HandlerFunc(rout.Handlers.SearchAnagrams)

	return router
}