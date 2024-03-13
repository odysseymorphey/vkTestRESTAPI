package main

import (
	"fmt"
	"net/http"
)

func main() {
	// TODO: Init config

	// TODO: Init logger

	//TODO: Init storage

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello", r.URL.String())
	})

	http.ListenAndServe(":8080", nil)
	fmt.Println("Penis")
}
