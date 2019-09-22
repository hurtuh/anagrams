package service

import (
	"encoding/json"
	"github.com/hurtuh/anagrams/domain"
	"log"
	"net/http"
)

type Service struct {
	Anagrams map[string][]string
}

func (serv *Service) SetAnagrams(w http.ResponseWriter, r *http.Request) {
	anagrams := domain.SliceAnagrams{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&anagrams)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	serv.saveAnagrams(&anagrams)
}


func (serv *Service) SearchAnagrams(w http.ResponseWriter, r *http.Request) {
	requet := domain.GetAnagrams{Anagram: domain.New(r.FormValue("word"))}
	res := serv.getAnagrams(&requet)
	bb, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, _ = w.Write(bb)
}