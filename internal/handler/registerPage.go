package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/smhmurat/golang-forum-frontend/models"
	"html/template"
	"net/http"
	"path/filepath"
)

type PageData struct {
	Message string
}

func ShowRegisterFormHandler(w http.ResponseWriter, r *http.Request) {
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

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := models.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Post data to register page
	res, err := http.Post("http://localhost:8080/api/v1/users", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if res.StatusCode == http.StatusCreated {
		//data := struct {
		//		//	Email    string
		//		//	Password string
		//		//	Username string
		//		//	Success  bool
		//		//}{
		//		//	Email:    email,
		//		//	Password: "",
		//		//	Username: username,
		//		//	Success:  true,
		//		//}
		http.Redirect(w, r, "/login.html", http.StatusSeeOther)
		return
	}
	fmt.Println("response Status:", res.Status)
}
