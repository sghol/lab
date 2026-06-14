package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "<h1>hello world</h1>")
		if err != nil {
			fmt.Println("Error :: ", err)
		}
	})

	_ = http.ListenAndServe(":8080", nil)
}
