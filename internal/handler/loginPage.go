package handler

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	layout := filepath.Join("web", "templates", "layout.html")
	navbar := filepath.Join("web", "templates", "navbar.html")
	login := filepath.Join("web", "templates", "login.html")
	search := filepath.Join("web", "templates", "search.html")
	tmpl, err := template.ParseFiles(layout, navbar, login, search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = tmpl.Execute(w, nil)
}
