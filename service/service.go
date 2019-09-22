package service

import (
	"github.com/hurtuh/anagrams/domain"
	"sort"
	"sync"
)

func (serv *Service) saveAnagrams(anagrams *domain.SliceAnagrams) {
	var wg sync.WaitGroup

	//Проходим по всем словам, что пришли в запросе
	for _, anagram := range anagrams.Anagrams {
		anagram := anagram
		wg.Add(1)
		go func() {
			defer wg.Done()

			//Сохраняем первоначальное слово
			value := anagram.String()
			sort.Sort(anagram)

			//Сохраняем после сортировки готовый ключ
			key := anagram.String()

			//Если мапа с данной аннаграмой уже существует, и если этого слова ещё нет, то сохраняем
			if _, ok := serv.Anagrams[key]; ok && !serv.checkDuplicate(value, key) {
				serv.Anagrams[key] = append(serv.Anagrams[key], value)

			} else if !ok {
				serv.Anagrams[key] = []string{value}
			}
		}()
		wg.Wait()
	}
}

func (serv *Service) getAnagrams(request *domain.GetAnagrams) *domain.SliceAnagrams {
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

func (serv *Service) checkDuplicate(request, key string) bool {
	for _, v := range serv.Anagrams[key] {
		if v == request {
			return true
		}
	}
	return false
}
