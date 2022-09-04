package main

import (
	"ascii-art-web/pkg/app"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.ViewHandler)
	mux.HandleFunc("/ascii-art", app.CreateHandler)
	mux.HandleFunc("/download", app.DownloadHandler)
	mux.Handle("/templates/css/", http.StripPrefix("/templates/css/", http.FileServer(http.Dir("templates/css"))))
	mux.Handle("/templates/img/", http.StripPrefix("/templates/img/", http.FileServer(http.Dir("templates/img"))))
	fmt.Printf("Starting server at post: 8087\nhttp://localhost:8087/\n")
	err := http.ListenAndServe("localhost:8086", mux)
	log.Fatal(err)
}
