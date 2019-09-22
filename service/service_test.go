package service

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/hurtuh/anagrams/domain"
)

var (
	anagrams1 = make(map[string][]string)
	handlers  = &Service{Anagrams: anagrams1}
)

// go test -v -race
func TestOne(t *testing.T) {

	anagrams2 := domain.SliceAnagrams{}
	err := json.Unmarshal([]byte(`{ "anagrams":["123","132","213","321","baba", "aabb","baba","baba","baba","baba","baba","aqwe", "кошка", "пёс", "aqwe", "кошка", "пёс","aqwe", "кошка", "пёс","aqwe", "кошка", "пёс","baab", "bbaa"]}`), &anagrams2)

	if err != nil {
		log.Fatal(err)
		return
	}
	handlers.saveAnagrams(&anagrams2)
	answer2 := []byte(`{"123":["123","132","213","321"],"aabb":["baba","aabb","baab","bbaa"],"aeqw":["aqwe"],"аккош":["кошка"],"псё":["пёс"]}`)
	answer1, err := json.Marshal(anagrams1)

	if string(answer2) != string(answer1) {
		t.Errorf("У нас вышло \n %s \n а должно быть\n %s\n", answer1, answer2)

	}
}

func TestTwo(t *testing.T) {

	requet := domain.GetAnagrams{Anagram: domain.New("123")}
	res := handlers.getAnagrams(&requet)
	bb, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		return
	}
	answer := `{"anagrams":["123","132","213","321"]}`
	if string(bb) != answer {
		t.Errorf("У нас вышло \n %s \n а должно быть\n %s\n", string(bb), answer)
	}

}
