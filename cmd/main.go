package main

import (
	"fmt"
	"log"
	"net/http"

	"export-file/internal/app"
)

func main() {
	http.HandleFunc("/", app.MainHandler)
	http.HandleFunc("/ascii-art", app.ArtHandler)
	http.HandleFunc("/download", app.DownloadHandler)
	http.Handle("/templates/css/", http.StripPrefix("/templates/css/", http.FileServer(http.Dir("templates/css"))))
	http.Handle("/templates/img/", http.StripPrefix("/templates/img/", http.FileServer(http.Dir("templates/img"))))
	fmt.Printf("Starting server at post: 8087\nhttp://localhost:8087/\n")
	log.Fatal(http.ListenAndServe(":8087", nil))
}
