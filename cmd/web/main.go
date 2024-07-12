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
	router.HandleFunc("/login.html", handler.LoginPage)
	router.HandleFunc("/register.html", handler.RegisterPage)

	// Starting the server
	log.Println("Starting server on :8081")
	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
