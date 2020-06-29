package main
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"encoding/json"
)

type SpeedsResponse struct{
	page string 'json:"page"'
	per_page int 'json:"per_page"'
	total int 'json:"total"'
	total_pages int 'json:"total_pages"'
	data Data 'json:"data"'
}

type Data struct{
	data []Info
}

type Info struct{
	id int 'json:"id"'
	timestamp int 'json:"timestamp"'
	

}

func Results(statusQuery string, page int){
	base := "https://jsonmock.hackerrank.com/api/iot_devices/search?"

	statusFilter := "status="+statusQuery
	pageFilter := "&page="+strconv.Itoa(page)

	url := base+statusFilter+pageFilter
	resp, err := http.Get(url)
	
	//fmt.Prinln(resp.Body)
	if err != nil{
		log.Fatal(err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	fmt.Println(string(contents))
	json.

}


func main(){
	Results("STOP", 2)
}