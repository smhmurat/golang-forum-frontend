package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/smhmurat/golang-forum-frontend/models"
	"io"
	//"github.com/smhmurat/golang-forum-frontend/models"
	"html/template"
	//"io"
	"net/http"
	"path/filepath"
	//"time"
)

func ShowLoginFormHandler(w http.ResponseWriter, r *http.Request) {
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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := models.User{
		Email:    email,
		Password: password,
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := http.Post("http://localhost:8082/auth/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if res.StatusCode == http.StatusOK {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(res.Body)
		body, err := io.ReadAll(res.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var loginResponse models.LoginResponse
		err = json.Unmarshal(body, &loginResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		userToken := loginResponse.Token
		http.SetCookie(w, &http.Cookie{
			Name:     "forum_session",
			Value:    userToken,
			HttpOnly: true,
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)

		return
	}

	//Hata durumunda yanıtı ve durum kodunu yazdırın
	fmt.Println("API yanıtı:", res.Status)
	body, err := io.ReadAll(res.Body)
	if err == nil {
		fmt.Println("API yanıt gövdesi:", string(body))
	}

	wrong := struct {
		Success bool
		Message string
	}{
		Success: false,
		Message: "Mail veya şifre yanlış",
	}
	fmt.Println(wrong)

	layout := filepath.Join("web", "templates", "layout.html")
	navbar := filepath.Join("web", "templates", "navbar.html")
	login := filepath.Join("web", "templates", "login.html")
	search := filepath.Join("web", "templates", "search.html")

	tmpl, err := template.ParseFiles(layout, navbar, login, search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = tmpl.Execute(w, wrong)
}
