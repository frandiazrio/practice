package main
import "fmt"

type Entry struct{
	phone string
	prefix string
}

func longestPrefix(numbers, prefixes []string)[]Entry{
	if numbers == nil || prefixes == nil || len(numbers) == 0 || len(prefixes) == 0{
		return []Entry{}
	}

	pref_mp := make(map[string]bool)

	for _, p := range prefixes{
		pref_mp[p]=false
	}
	res := []Entry{}
	for _, number := range numbers{
		n := number
		for len(number) > 1{
			t, ok := pref_mp[number]

			if ok && !t {
				res = append(res, Entry{n, number})
				pref_mp[number]= true
				break
			}else{
				number = number[: len(number)-1] 
			}
		}
	}

	return res
}

func main(){
	
	prefixes1 := []string{"+1415123", "+1415000", "+1412", "+1510", "+1", "+44"}
	numbers1 := []string{"+14151234567", "+99999999999", "+14121234567", "+19999999999"}
	fmt.Println(longestPrefix(numbers1, prefixes1))
	
	prefixes2 := []string{"+1415123"}
	numbers2 := []string{"+99999999999"}
	fmt.Println(longestPrefix(numbers2, prefixes2))



	prefixes3 := []string{"+1415123"}
	numbers3 := []string{"+14151234567"}
	fmt.Println(longestPrefix(numbers3, prefixes3))
}