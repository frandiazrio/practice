package main

import (
	"fmt"
	"strconv"
)

type Set struct {
	items map[string]struct{}
}

func (s *Set) Add(item string) {
	s.items[item] = struct{}{}
}

func (s *Set) Contains(item string) bool {
	_, exists := s.items[item]
	return exists
}

func (s *Set) Delete(item string) {
	delete(s.items, item)
}

func main() {

	item_set := Set{}

	for i := 0; i < 10; i++ {
		item_set.Add(strconv.Itoa(i))
	}

	fmt.Println(item_set)

}
