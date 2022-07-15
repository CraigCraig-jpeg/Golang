package main

import (
	"net/http"
	"fmt"
)

type server struct {
	
}

func (m server) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w , "hello")
}

func main() {
	var serve server
	http.ListenAndServe(":8080", serve)
}