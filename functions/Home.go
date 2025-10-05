package functions

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		RenderError(w, "Page not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		RenderError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		RenderError(w, "Please try later", http.StatusInternalServerError)
		fmt.Println("Failed to parse index template", err)
		return
	}

	var artists []Artist
	var buf bytes.Buffer

	err2 := FetchJson(url, &artists)
	if err2 != nil {
		RenderError(w, "Please try later", http.StatusInternalServerError)
		fmt.Println(err2)
		return
	}

	err3 := tmpl.Execute(&buf, &artists)
	if err3 != nil {
		RenderError(w, "Please try later", http.StatusInternalServerError)
		fmt.Println("Failed to execute index template", err3)
		return
	}

	_, err4 := buf.WriteTo(w)
	if err4 != nil {
		RenderError(w, "Please try later", http.StatusInternalServerError)
		fmt.Println("Failed to write in the buffer")
		return
	}
}
