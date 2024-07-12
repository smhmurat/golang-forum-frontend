package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	layout := filepath.Join("web", "templates", "layout.html")
	index := filepath.Join("web", "templates", "index.html")
	navbar := filepath.Join("web", "templates", "navbar.html")
	search := filepath.Join("web", "templates", "search.html")
	post := filepath.Join("web", "templates", "posts.html")

	tmpl, err := template.ParseFiles(layout, index, navbar, search, post)
	if err != nil {
		fmt.Println("Error parsing templates:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
