package app

import (
	"net/http"
	"strconv"
	"text/template"

	"export-file/internal/service"
)

var artist service.Artist

func MainHandler(writer http.ResponseWriter, request *http.Request) {
	artist.UpdateArtist()
	if request.URL.Path != "/" {
		ErrorHandler(writer, "Page is Not Found", http.StatusNotFound)
		return
	}

	if request.Method != http.MethodGet {
		ErrorHandler(writer, "Method is not Allowed", http.StatusMethodNotAllowed)
		return
	}
	html, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = html.Execute(writer, nil)
	if err != nil {
		ErrorHandler(writer, "Page is Not Found", http.StatusNotFound)
		return
	}
}

func ArtHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("templates/answer.html")
	if err != nil {
		ErrorHandler(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if request.Method != http.MethodPost {
		ErrorHandler(writer, "Method is not Allowed", http.StatusMethodNotAllowed)
		return
	}

	artist.UpdateText(request.FormValue("text"))
	if artist.ErrMessage != "" {
		ErrorHandler(writer, "Bad Request. "+artist.ErrMessage, http.StatusBadRequest)
		return
	}
	artist.UpdateBanner(request.FormValue("banner"))
	if artist.ErrMessage != "" {
		ErrorHandler(writer, "Bad Request. "+artist.ErrMessage, http.StatusBadRequest)
		return
	}
	artist.Draw()
	if artist.HasError() {
		ErrorHandler(writer, "Internal Server Error. "+artist.ErrMessage, http.StatusInternalServerError)
		return
	}

	err = html.Execute(writer, artist)
	if err != nil {
		http.Error(writer, "404: Not Found", 404)
		return
	}
}

func DownloadHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		ErrorHandler(writer, "Method is not Allowed", http.StatusMethodNotAllowed)
		return
	}
	writer.Header().Add("Content-Length", strconv.Itoa(len(artist.Art)))
	writer.Header().Add("Content-Language", "en")
	writer.Header().Add("Content-Disposition", "attachment; filename=ascii-art.txt")
	writer.Header().Add("Content-Type", "text/plain; charset=utf-8")
	writer.Write([]byte(artist.Art))
}
