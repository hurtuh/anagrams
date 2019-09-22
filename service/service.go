package service

import (
	"fmt"
	"github.com/anagrams/domain"
	"sort"
)

func (serv *Service) SaveAnagrams(anagrams *domain.SliceAnagrams) {
	for _, anagram := range anagrams.Anagrams {
		value := anagram.String()
		//fmt.Println(anagram.String())
		sort.Sort(anagram)
		//fmt.Println(anagram.String())
		if _, ok := serv.Anagrams[anagram.String()]; ok {
			serv.Anagrams[anagram.String()] = append(serv.Anagrams[anagram.String()], value)
			continue
		}
		serv.Anagrams[anagram.String()] = []string{value}
	}
}

func (serv *Service) GetAnagrams(anaram *domain.GetAnagrams) *domain.SliceAnagrams {
	for _, v := range serv.Anagrams {
		fmt.Println(v)
	}
	return nil
}
