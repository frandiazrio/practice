package main
import "fmt"
func longestPrefix(numbers, prefixes []string)[]string{
	if numbers == nil || prefixes == nil || len(numbers) == 0 || len(prefixes) == 0{
		return []string{}
	}

	pref_mp := make(map[string]struct{})

	for _, p := range prefixes{
		pref_mp[p]=struct{}{}
	}
	res := []string{}
	for _, number := range numbers{
		for len(number) > 0{
			_, ok := pref_mp[number]

			if ok {
				res = append(res, number)
				break
			}else{
				number = number[: len(number)-1] 
			}
		}
	}

	return res
}

func main(){
}