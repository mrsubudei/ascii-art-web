package app

import (
	"ascii-art-web/pkg/ascii"
	"net/http"
	"text/template"
)

type Asc struct {
	Str      string
	Orig     string
	StatCode int
}

var Ascii Asc

func ViewHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		Errors(writer, "Page is Not Found", http.StatusNotFound)
		return
	}

	if request.Method != http.MethodGet {
		Errors(writer, "Method is not Allowed", http.StatusMethodNotAllowed)
		return
	}

	html, err := template.ParseFiles("templates/index.html")
	if err != nil {
		Errors(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = html.Execute(writer, nil)
	if err != nil {
		Errors(writer, "Page is Not Found", http.StatusNotFound)
		return
	}
}

func CreateHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("templates/answer.html")
	if err != nil {
		Errors(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if request.Method != http.MethodPost {
		Errors(writer, "Method is not Allowed", http.StatusMethodNotAllowed)
		return
	}

	req := request.FormValue("req")
	font := request.FormValue("font")

	if req == "" {
		Ascii.StatCode = 400
		Errors(writer, "Bad Request. Write something", http.StatusBadRequest)
		return
	}

	if font == "" {
		Ascii.StatCode = 400
		Errors(writer, "Bad Request. Choose any font", http.StatusBadRequest)
		return
	}

	ans, statCode := ascii.ConvertAscii(req, font)
	Ascii.StatCode = statCode
	if Ascii.StatCode == 200 {
		Ascii.Str = ans
		Ascii.Orig = req
	} else {
		Errors(writer, ans, statCode)
		return
	}

	err = html.Execute(writer, Ascii)
	if err != nil {
		http.Error(writer, "404: Not Found", 404)
		return
	}
}

// func DownloadHandler(writer http.ResponseWriter, request *http.Request) {
// 	if request.Method != http.MethodGet {
// 		Errors(writer, "Method is not Allowed", http.StatusMethodNotAllowed)
// 		return
// 	}
// 	writer.Header().Add("Content-Length", strconv.Itoa(len(Ascii.Str)))
// 	writer.Header().Add("Content-Language", "en")
// 	writer.Header().Add("Content-Disposition", "attachment; filename=ascii-art.txt")
// 	writer.Header().Add("Content-Type", "text/plain; charset=utf-8")
// 	writer.Write([]byte(Ascii.Str))
// }

func Errors(writer http.ResponseWriter, er string, code int) {
	writer.WriteHeader(code)
	Ascii.Str = er
	Ascii.StatCode = code
	html, err := template.ParseFiles("templates/errors.html")
	if err != nil {
		http.Error(writer, "500: Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = html.Execute(writer, Ascii)
	if err != nil {
		http.Error(writer, "404: Not Found", 404)
		return
	}
}
