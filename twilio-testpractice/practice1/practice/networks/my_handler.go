package main

import (
	"fmt"
	"net/http"
)

type RandomType int

func (r RandomType) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(res, "Hello from Handler")
}

func main() {
	var handler RandomType
	http.ListenAndServe(":8080", handler)
}
