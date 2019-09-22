package service

import (
	"encoding/json"
	"fmt"
	"github.com/anagrams/domain"
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
		fmt.Println(err)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(len(anagrams.Anagrams[0].String()), "THIS")
	serv.SaveAnagrams(&anagrams)
}


func (serv *Service) SearchAnagrams(w http.ResponseWriter, r *http.Request) {
	serv.GetAnagrams(&domain.GetAnagrams{})
}