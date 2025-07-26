package handlers

import (
	"ascii-art-web/core"
	"log"
	"net/http"
	"text/template"
)

func isAscii(str string) bool {
	for _, e := range str {
		if e > 127 {
			return false
		}
	}
	return true
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		notFoundHandler(w)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Template error: %v", err)
		internalServerError(w)
		return
	}
	title, err := core.ProcessInput("ASCII Art Generator", "standard")
	if err != nil {
		internalServerError(w)
		return
	}

	if err := tmpl.Execute(w, title); err != nil {
		log.Printf("Template execution error: %v", err)
		internalServerError(w)
		return
	}
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		notAllowedMethod(w)
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Printf("Form parse error: %v", err)
		badRequestHandler(w, "Form parse error")
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" {
		badRequestHandler(w, "You should submit a text")
		return
	}

	if !isAscii(text) {
		badRequestHandler(w, "Only ASCII characters are allowed")
		return
	}
	title, err := core.ProcessInput("ASCII Art Result", "standard")
	if err != nil {
		internalServerError(w)
		return
	}

	result, err := core.ProcessInput(text, banner)
	if err != nil {
		log.Printf("Processing error: %v", err)
		internalServerError(w)
		return
	}

	tmpl, err := template.ParseFiles("templates/result.html")
	if err != nil {
		log.Printf("Template error: %v", err)
		internalServerError(w)
		return
	}

	data := struct {
		Text   string
		Banner string
		Result string
		Title  string
	}{
		Text:   text,
		Banner: banner,
		Result: result,
		Title:  title,
	}

	if err := tmpl.Execute(w, data); err != nil {
		internalServerError(w)
	}
}

func notFoundHandler(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)

	tmpl, err := template.ParseFiles("errors/404.html")
	if err != nil {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}

func badRequestHandler(w http.ResponseWriter, data string) {
	w.WriteHeader(http.StatusBadRequest)

	tmpl, err := template.ParseFiles("errors/400.html")
	if err != nil {
		http.Error(w, "400 Not Found", http.StatusNotFound)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}

func notAllowedMethod(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)

	tmpl, err := template.ParseFiles("errors/405.html")
	if err != nil {
		http.Error(w, "405 Not Found", http.StatusNotFound)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}

func internalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)

	tmpl, err := template.ParseFiles("errors/500.html")
	if err != nil {
		http.Error(w, "500 Not Found", http.StatusNotFound)
		return
	}

	err = tmpl.Execute(w, http.StatusInternalServerError)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}

func ExportFile(w http.ResponseWriter, r *http.Request) {

}
