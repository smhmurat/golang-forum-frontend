package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/smhmurat/golang-forum-frontend/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

type PageData struct {
	Message string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading.env file")
	}
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
	res, err := http.Post("http://localhost:8082/auth/signup", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if res.StatusCode == http.StatusCreated {
		//data := struct {
		//			Email    string
		//			Password string
		//			Username string
		//			Success  bool
		//		}{
		//			Email:    email,
		//			Password: "",
		//			Username: username,
		//			Success:  true,
		//		}
		http.Redirect(w, r, "/login.html", http.StatusSeeOther)
		return
	}
	fmt.Println("response Status:", res.Status)
}

func RegisterWithGoogleHandler(w http.ResponseWriter, r *http.Request) {
	res, err := http.Post("http://localhost:8082/auth/google/signup", "application/json", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if res.StatusCode != http.StatusBadRequest {
		expiration := time.Now().Add(24 * time.Hour)
		cookie := http.Cookie{Name: "session", Value: "true", Expires: expiration}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Println("User logged in successfully")
		return
	}

}
