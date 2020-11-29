package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"
	fmt.Fprintf(w, msg)

}
