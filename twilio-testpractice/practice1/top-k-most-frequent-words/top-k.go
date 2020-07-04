package main

import "fmt"
import "sort"

type Entry struct{
	word string
	count int
}

func topK(words []string, k int)[]string{
	if len(words) < 1 || k < 0 {
		return nil
	}

	mp := make(map[string]int)

	for _, w := range words{
		mp[w]++
	}

	tmp := []Entry{}

	for k,v := range mp{
		tmp = append(tmp, Entry{k,v})
	}

	sort.Slice(tmp, func(i, j int)bool{
		return tmp[i].count > tmp[j].count 
	})

	res := []string{}

	for i:=0; i< k; i++{
		res = append(res, tmp[i].word)
	}
	return res

}

func main(){

	fmt.Println(topK([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))

}