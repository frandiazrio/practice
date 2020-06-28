package main


import (
	"fmt"
	"strings"
)


func missing(s,t string)[]string{
	mp := make(map[string]struct{})


	res := []string{}
	for _,v := range strings.Split(t, " "){
		mp[v] = struct{}{}
	}

	for _,v := range strings.Split(s, " "){
		_,ok := mp[v]
		if !ok{
			res = append(res, v)
		}
	}


	return res
}






func main(){
	fmt.Println(missing("I like cheese", "I like"))
}