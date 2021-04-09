package main

import (
	"html/template"
	"log"
	"net/http"
)

const (
	Host = "localhost"
	Port = "8080"
)

type HomePage struct {
  Title      string
	Name       string
	College    string
	RollNumber int
}

func Home(w http.ResponseWriter, r *http.Request) {
	homepage := HomePage{
	  Title:      "Learning Golang",
		Name:       "Gourab",
		College:    "Golang Docs",
		RollNumber: 103,
	}
	parsedTemplate, _ := template.ParseFiles("template/index.html")
	err := parsedTemplate.Execute(w, homepage)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func main() {
	http.HandleFunc("/", Home)
	err := http.ListenAndServe(Host+":"+Port, nil)
	log.Fatal("server started at 8080 port")
	if err != nil {
		log.Fatal("Error Starting the HTTP Server :", err)
		return
	}
}