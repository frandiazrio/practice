package main

import (
	"fmt"
	"strings"
)

func missing(s, t string) []string {
	mp := make(map[string]struct{})

	res := []string{}
	for _, v := range strings.Split(t, " ") {
		mp[v] = struct{}{}
	}

	for _, v := range strings.Split(s, " ") {

		_, ok := mp[v]
		if !ok && v != "" {
			res = append(res, v)
		}
	}
	
	return res
}

func missing2(s, t string) []string {
	set := make(map[string]struct{})
	res := []string{}
	for _, v := range strings.Split(t, " ") {
		set[v] = struct{}{}
	}

	for v, _ := range set {
		//fmt.Println(v,strings.Contains(t,v))
		if !strings.Contains(s, v) {
			res = append(res, v)
		}
	}

	return res
}

func main() {
	fmt.Println(missing("I am using    HackerRank to improve programming", "am HackerRank to improve"))
	//fmt.Println(missing2("I like cheese", "I like"))
}
