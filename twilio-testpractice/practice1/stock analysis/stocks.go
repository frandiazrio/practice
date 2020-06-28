package main

import "strings"
import "fmt"
func eliminate(arr []string){
	mp := make(map[string]string)
	for _, source := range arr{
		for _, v := range strings.Split(source, " "){
			update(v, mp)
		}
	}

	for _,v := range mp{
		fmt.Println(v)
	}
}

func update(v string, mp map[string]string){
	tmp := strings.Split(v,":")
	key := tmp[0]
	value := tmp[1]
	mp[key] = value
}

func main(){
	input := []string{"P1:a P3:b P5:x","P1:b P2:q P5:x"}
	eliminate(input)
}