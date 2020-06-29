package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	
)

func hostsRequests(logfile string) {
	f, err := os.Open(logfile)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	mp := make(map[string]int) // host name to number of requests
	for scanner.Scan() {
		line := scanner.Text()
		
		tmp := strings.Split(line, "- -")
		url := tmp[0]
	

		mp[url]++

	}

	file, err := os.Create("records_filename")
	defer file.Close()
	if err != nil{
		log.Fatal(err)
	}

	
	for hosts,num := range mp{
		file.WriteString(hosts+" "+strconv.Itoa(num)+"\n")
	}

	fmt.Println(mp)

}

func main() {
	hostsRequests("logfile.txt")
}
