package handler

import (
	"fmt"
	"github.com/smhmurat/golang-forum-frontend/models"
	"html/template"
	"net/http"
	"path/filepath"
)

func getUserSession(r *http.Request) models.UserSession {
	session, err := r.Cookie("forum_session")
	if err != nil {
		return models.UserSession{LoggedIn: false}
	}
	userSession := models.UserSession{LoggedIn: false}
	if session != nil {
		userSession.LoggedIn = true
	}
	return userSession
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

	user := getUserSession(r)

	err = tmpl.Execute(w, user)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
