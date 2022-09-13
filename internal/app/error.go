package app

import (
	"net/http"
	"text/template"
)

func ErrorHandler(writer http.ResponseWriter, er string, code int) {
	writer.WriteHeader(code)
	artist.ErrMessage = er
	artist.StatusCode = code
	html, err := template.ParseFiles("templates/errors.html")
	if err != nil {
		http.Error(writer, "500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = html.Execute(writer, artist)
	if err != nil {
		http.Error(writer, "404: Not Found", 404)
		return
	}
}
