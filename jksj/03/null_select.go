package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello, gohercon sg")
	})
	go func() {}()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	select {}
}
