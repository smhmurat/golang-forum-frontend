package main

import (
	"log"
	"os"
)

type application struct {
	appName string
	server  server
	debug   bool
	errLog  *log.Logger
	infoLog *log.Logger
}

type server struct {
	host string
	port string
	url  string
}

func main() {
	server := server{
		host: "localhost",
		port: "8080",
		url:  "http://localhost:8080",
	}
	app := &application{
		server:  server,
		appName: "Golang Forum Web Application",
		debug:   true,
		errLog:  log.New(os.Stderr, "ERROR: ", log.Ltime|log.Ldate|log.Llongfile),
		infoLog: log.New(os.Stdout, "INFO: ", log.Ltime|log.Ldate|log.Llongfile),
	}

	if err := app.listenAndServer(); err != nil {
		log.Fatal(err)
	}
}
