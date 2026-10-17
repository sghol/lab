package main

import (
	"fmt"
	"github.com/sghol/go-htmx/database"
	"html/template"
	"net/http"
)

// template config
var templates = template.Must(template.ParseFiles(
	"templates/index.html",
	"templates/greet.html",
	"templates/students.html",
))

type handler struct {
}

// home page
func (h *handler) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	templates.ExecuteTemplate(w, "index.html", map[string]string{"Name": "Saeed Gholami"})
}

// greet page
func (h *handler) Greet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	templates.ExecuteTemplate(w, "greet.html", nil)
}

func (h *handler) Students(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	students, err := database.GetStudents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templates.ExecuteTemplate(w, "students.html", students)
}

// entry point
func main() {
	database.Init()
	h := handler{}

	// routes
	http.HandleFunc("/", h.Index)
	http.HandleFunc("/greet", h.Greet)
	http.HandleFunc("/students", h.Students)

	// serve
	fmt.Println("Listen on :8080")
	http.ListenAndServe(":8080", nil)
}
