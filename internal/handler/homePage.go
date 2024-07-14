package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type User struct {
	LoggedIn bool
}

func GetUser(r *http.Request) User {
	session, _ := r.Cookie("session")
	user := User{LoggedIn: false}
	if session != nil && session.Value == "true" {
		user.LoggedIn = true
	}
	return user
}

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

	user := GetUser(r)

	err = tmpl.Execute(w, user)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
