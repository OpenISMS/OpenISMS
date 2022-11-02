package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var tmpl *template.Template

//go:embed templates/*
var templates embed.FS

func main() {
	mux := http.NewServeMux()
	tmpl, _ = template.ParseFS(templates, "templates/*")

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", dashboard)
	mux.HandleFunc("/people", people)
	mux.HandleFunc("/devices", devices)
	mux.HandleFunc("/suppliers", suppliers)
	mux.HandleFunc("/keys", keys)
	mux.HandleFunc("/apps", apps)
	mux.HandleFunc("/context", context)

	listen := getEnv("LISTEN", "localhost:8085")
	fmt.Printf("Go To http://%s/\n", listen)
	log.Fatal(http.ListenAndServe(listen, mux))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func keys(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "keys.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func apps(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "apps.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func dashboard(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "dashboard.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func context(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "context.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func devices(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "devices.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}
