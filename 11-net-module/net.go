package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")

	}
	http.HandleFunc("/bar", h1)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
