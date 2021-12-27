package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fprintf, err := fmt.Fprintf(w, "ULR.Path = %q\n", r.URL.Path)
	if err != nil {
		return
	}
	fmt.Printf("Bytes written: %v\n", fprintf)
}
