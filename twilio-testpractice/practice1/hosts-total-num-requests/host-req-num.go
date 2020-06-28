package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	
)

func hostsRequests(logfile string)map[string]int {
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
		fields := strings.Split(url,".")
		hostname := fields[len(fields)-2]
		if err != nil{
			log.Fatal(err)
		}

		mp[hostname]++

	}

	return mp

}

func main() {
	fmt.Println(hostsRequests("logfile.txt"))
}
