package service

import (
	"github.com/anagrams/domain"
	"sort"
)

func (serv *Service) SaveAnagrams(anagrams *domain.SliceAnagrams) {
	for _, anagram := range anagrams.Anagrams {
		value := anagram.String()
		sort.Sort(anagram)
		if _, ok := serv.Anagrams[anagram.String()]; ok {
			serv.Anagrams[anagram.String()] = append(serv.Anagrams[anagram.String()], value)
			continue
		}
		serv.Anagrams[anagram.String()] = []string{value}
	}
}

func (serv *Service) GetAnagrams(request *domain.GetAnagrams) *domain.SliceAnagrams {
	sort.Sort(request.Anagram)

	if _, ok := serv.Anagrams[request.Anagram.String()]; ok {
		response := &domain.SliceAnagrams{}

		for _, value := range serv.Anagrams[request.Anagram.String()] {
			response.Anagrams = append(response.Anagrams, domain.New(value))
		}

		return response
	}
	return nil
}
