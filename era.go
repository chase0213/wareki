package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/dghubble/trie"
)

type EraSearcher struct {
	Eras []Era
	Trie trie.Trier
}

type Era struct {
	Name       string `json:"name"`
	Yomi       string `json:"yomi"`
	BeginYear  int    `json:"begin_year"`
	BeginMonth int    `json:"begin_month"`
	BeginDay   int    `json:"begin_day"`
	EndYear    int    `json:"end_year"`
	EndMonth   int    `json:"end_month"`
	EndDay     int    `json:"end_day"`
}

type dataFromJson struct {
	Era     string `json:"era"`
	Yomi    string `json:"yomi"`
	BeginAt string `json:"begin_at"`
	EndAt   string `json:"end_at"`
}

func (es *EraSearcher) Search(query string) (Era, error) {
	var era Era
	jsonEra := fmt.Sprintf("%s", es.Trie.Get(query))
	if err := json.Unmarshal([]byte(jsonEra), &era); err != nil {
		return Era{}, err
	}

	return era, nil
}

func NewEraSearcher() (*EraSearcher, error) {
	eras, err := LoadEras()
	if err != nil {
		return nil, err
	}

	trie, err := constructTrieFromEras(eras)

	es := EraSearcher{
		Eras: eras,
		Trie: trie,
	}

	return &es, nil
}

func LoadEras() ([]Era, error) {
	bytes, err := ioutil.ReadFile("eras.json")
	if err != nil {
		return nil, err
	}

	var data []dataFromJson
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	eras := make([]Era, 0, len(data))
	for _, d := range data {
		era, err := parseEra(d)
		if err != nil {
			return []Era{}, err
		}

		eras = append(eras, era)
	}

	return eras, nil
}

func constructTrieFromEras(eras []Era) (*trie.RuneTrie, error) {
	trie := trie.NewRuneTrie()
	for _, era := range eras {
		jsonEra, _ := json.Marshal(era)
		trie.Put(era.Name, jsonEra)
	}
	return trie, nil
}

func parseEra(data dataFromJson) (Era, error) {
	era := Era{}
	era.Name = data.Era
	era.Yomi = data.Yomi

	era.BeginYear = -1
	era.BeginMonth = -1
	era.BeginDay = -1
	era.EndYear = -1
	era.EndMonth = -1
	era.EndDay = -1

	var err error
	beginAtSlice := strings.Split(data.BeginAt, "-")
	beginYear, err := strconv.Atoi(beginAtSlice[0])
	if err != nil {
		return era, nil
	}
	beginMonth, err := strconv.Atoi(beginAtSlice[1])
	if err != nil {
		return era, nil
	}
	beginDay, err := strconv.Atoi(beginAtSlice[2])
	if err != nil {
		return era, nil
	}

	era.BeginYear = beginYear
	era.BeginMonth = beginMonth
	era.BeginDay = beginDay

	endAtSlice := strings.Split(data.EndAt, "-")
	endYear, err := strconv.Atoi(endAtSlice[0])
	if err != nil {
		return era, nil
	}
	endMonth, err := strconv.Atoi(endAtSlice[1])
	if err != nil {
		return era, nil
	}
	endDay, err := strconv.Atoi(endAtSlice[2])
	if err != nil {
		return era, nil
	}

	era.EndYear = endYear
	era.EndMonth = endMonth
	era.EndDay = endDay

	return era, nil
}
