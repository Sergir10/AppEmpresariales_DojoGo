package main

import (
	"fmt"
	"net/http" //package for http based web programs
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside handler")
	fmt.Fprintf(w, "Hello world from my Go program!")
}

func main() {
	http.HandleFunc("/", handler)              // redirect all urls to the handler function
	http.ListenAndServe("localhost:9999", nil) // listen for connections at port 9999 on the local machine
}
