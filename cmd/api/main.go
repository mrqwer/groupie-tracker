package main

import (
	"fmt"
	config "groupie-tracker/internal"
	"groupie-tracker/internal/web"
	"html/template"
	"log"
	"net/http"
	"time"
)

var (
	mux *http.ServeMux
	srv *http.Server
	tpl *template.Template
)

// func init() {
// 	fmt.Println("Starting up.")
// 	tpl, err := template.ParseGlob("ui/templates/*.html")
// 	if err != nil {
// 		log.Fatal("Error loading templates:" + err.Error())
// 	}
// 	routes.Tpl = tpl
// }

func main() {
	web.Client = &http.Client{Timeout: 10 * time.Second}
	mux = http.NewServeMux()

	mux.HandleFunc("/", mainHandler)
	mux.HandleFunc("/artist/", detailHandler)
	fs := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	tpl, _ = tpl.ParseGlob("ui/templates/*.html")

	srv = &http.Server{
		Addr:         fmt.Sprintf(":%s", config.Port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("starting server on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
