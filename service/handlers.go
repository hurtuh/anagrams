package service

import (
	"encoding/json"
	"github.com/anagrams/domain"
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
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	serv.SaveAnagrams(&anagrams)
}


func (serv *Service) SearchAnagrams(w http.ResponseWriter, r *http.Request) {
	requet := domain.GetAnagrams{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&requet)
	if err != nil {
		log.Println(err)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	res := serv.GetAnagrams(&requet)
	bb, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, _ = w.Write(bb)
}