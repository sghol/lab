package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomeHandler)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"hello": "world",
	}

	js, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
}
