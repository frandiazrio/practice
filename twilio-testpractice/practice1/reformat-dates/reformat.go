package main

import (
	"fmt"
	"strconv"
	"strings"
)


func ReformatDate(date string)(string) {
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep" , "Oct" , "Nov", "Dec"}
	mp :=map[string]string{
		"1st":"01",
		"2nd":"02",
		"3rd":"03",
		"32st" :"31",
	}


	for i :=4; i<10;i++{
		mp[strconv.Itoa(i)+"th"]="0"+strconv.Itoa(i)
	}

	for i :=10; i<31;i++{
		mp[strconv.Itoa(i)+"th"]=strconv.Itoa(i)
	}

	

	

	split := strings.Split(date, " ")
	if len(split) != 3{
		return ""
	}

	day := split[0]
	month := split[1]
	year := split[2]

	newDay, ok := mp[day]
	if !ok {
		return ""
	}

	newMonth := monthNumber(months, month)



	return year+"-"+newMonth+"-"+newDay

}

func monthNumber(months []string, month string)(string){
	for i, m := range months{
		if m == month{
			if i+1 < 9{
				return "0"+strconv.Itoa(i+1)
			}
			return strconv.Itoa(i+1)
		}
	}
	return ""
}

func main(){
	fmt.Println(ReformatDate("1st Mar 1984"))
	fmt.Println(ReformatDate("2nd Feb 2013"))
	fmt.Println(ReformatDate("4th Apr 1990"))
}