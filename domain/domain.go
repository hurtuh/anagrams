package domain

import (
	"encoding/json"
	"fmt"
)
//{ "anagrams":["baba", "aqwe", "кошка", "пёс"]}

type SliceAnagrams struct {
	Anagrams []Anagram `json:"anagrams"`
}

type GetAnagrams struct {
	Anagram Anagram `json:"anagram"`
}

type Anagram struct {
	sliceRune []rune
}

func (a Anagram) UnmarshalJSON(b []byte) error {
	tmp := ""
	var err error
	if err = json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	a.sliceRune =  []rune(tmp)
	fmt.Println(a.sliceRune)
	return nil
}

func (a Anagram) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(a.sliceRune))
}

func (a Anagram) Len() int {
	return len(a.sliceRune)
}

func (a Anagram) Swap(i, j int) {
	a.sliceRune[i], a.sliceRune[j] = a.sliceRune[j], a.sliceRune[i]
}

func (a Anagram) Less(i, j int) bool {
	return a.sliceRune[i] < a.sliceRune[j]
}

func (a Anagram) String() string {
	return string(a.sliceRune)
}