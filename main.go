package main

import (
	"fmt"
	"github.com/anagrams/routes"
	"github.com/anagrams/service"
	"log"
	"net/http"
)

func main() {

	anagrams := make(map[string][]string)

	handlers := &service.Service{Anagrams:anagrams}
	routes := routes.Routers{Handlers:handlers}

	router := routes.MakeHTTPHandlers()

	fmt.Println("Start server")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}