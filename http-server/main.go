package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func main() {
	store := InMemoryPlayerStore{
		map[string]int{},
	}
	handler := &PlayerServer{&store}
	log.Fatal(http.ListenAndServe(":3000", handler))
}
