package handler

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	layout := filepath.Join("web", "templates", "layout.html")
	navbar := filepath.Join("web", "templates", "navbar.html")
	register := filepath.Join("web", "templates", "register.html")
	search := filepath.Join("web", "templates", "search.html")

	tmpl, err := template.ParseFiles(layout, navbar, register, search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = tmpl.Execute(w, nil)
}
