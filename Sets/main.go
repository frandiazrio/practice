package main

import "fmt"

import "errors"

type Set struct {
	mp map[interface{}]struct{}
}

func InitSet(strs []string) *Set {
	mp := make(map[interface{}]struct{})
	set := &Set{mp: mp}
	for _, v := range strs {
		set.Add(v)
	}

	return set
}

func (s *Set) Add(str interface{}) {
	s.mp[str] = struct{}{}
}

func (s *Set) Remove(str interface{}) (interface{}, error) {
	_, ok := s.mp[str]

	if ok {
		delete(s.mp, str)
		return str, nil
	}

	return "", errors.New("Element not found in set")

}

func (s *Set) Union(st *Set) {

	for v, _ := range st.mp {
		s.Add(v)
	}

}
func PowerSet(st *Set) *Set{
	if len(st.mp) == 0 {

		return &Set{}

	var take interface{}
	for v, _ := range st.mp {
		take = v
		break
	}

	fmt.Println(st)
	st.Remove(take)
	res :=
	for _, set := range PowerSet(st) {
		res = append(res, set)
		set.Add(take)
		res = append(res, set)
	}

	return res

}
func main() {

	set1 := InitSet([]string{"1", "2", "3", "4"})
	set2 := InitSet([]string{"1", "2", "4", "2", "5", "1"})

	set1.Union(set2)

	fmt.Println(set1)

	fmt.Println(PowerSet(set1))
}
