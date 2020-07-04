package main

import (
	"fmt"
	"strings"
	"unicode"
)

func letNum(letter string)string{ 
	if strings.Contains("ABC", letter){
		return "2"	
	}
	if strings.Contains("DEF", letter){
		return "3"
	} 
	if strings.Contains("GHI", letter){
		return "4"
	}
	if strings.Contains("JKL", letter){
		return "5"
	}
	if strings.Contains("MNO", letter){
		return "6"
	}
	if strings.Contains("PQRS", letter){
		return "7"
	}
	if strings.Contains("TUV", letter){
		return "8"
	}
	if strings.Contains("WXYZ", letter){
		return "9"
	}
	return ""

}

func let_nums(str string)string{
	res := ""
	for _,c := range str {
		if unicode.IsLetter(c){
			res = res + letNum(string(c))
		}else{
			res = res+string(c)
		}
	}
	return res
}

func getCodes(codes[]string)[]string{
	res := []string{}
	for _,c := range codes{
		res = append(res, let_nums(strings.ToUpper(c)))
	}
	return res
}

func Contains(codes, numbers []string)[][]string{
	res := [][]string{}

	for _, n := range numbers{

		for _,c := range codes{
			if strings.Contains(n,c){
				res = append(res, []string{n,c})
			}
		}
		
	}

	return res
}

func main(){
	codes := []string{"PIZZA", "WATER", "FBI"}
	numbers := []string{"7878787774992", "82828292837", "82828282324"}
	fmt.Println(getCodes(codes))

	fmt.Println(Contains(getCodes(codes), numbers))
	


}