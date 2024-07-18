package main

import (
	"github.com/gorilla/mux"
	"github.com/smhmurat/golang-forum-frontend/internal/handler"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	fileServer := http.FileServer(http.Dir("./web/static/"))
	router.PathPrefix("/web/static/").Handler(http.StripPrefix("/web/static/", fileServer))

	router.HandleFunc("/", handler.HomePage)
	router.HandleFunc("/register.html", handler.ShowRegisterFormHandler)
	router.HandleFunc("/register", handler.RegisterHandler)
	router.HandleFunc("/login.html", handler.ShowLoginFormHandler)
	router.HandleFunc("/login", handler.LoginHandler)
	router.HandleFunc("/logout", handler.LogoutHandler)

	// Starting the server
	log.Println("Starting server on :8081")
	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
