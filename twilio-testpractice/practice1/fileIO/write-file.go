package main

import (
	"log"
	"os"
)


func main(){
	file, err := os.Create("tmp1.txt")
	defer file.Close()
	if err != nil{
		log.Fatal(err)
	}

	linesToWrite := []string{"This is an example", "to show how", "to write to a file", "line by line."}

	for _,line := range linesToWrite{
		file.WriteString(line)
	}


}