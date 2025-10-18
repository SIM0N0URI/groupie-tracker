package functions

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// About handles the display of a detailed artist page based on the artist ID.
// It validates the URL, fetches the artist s full data, and renders the corresponding HTML template.
func About(w http.ResponseWriter, r *http.Request) {
	ID := strings.TrimPrefix(r.URL.Path, "/artists/")

	if ID == "" {
		RenderError(w, "Page not found", http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(ID)
	if err != nil || id < 1 || id > 52 {
		RenderError(w, "Page not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		RenderError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		RenderError(w, "Please try later", http.StatusInternalServerError)
		fmt.Println("Failed to parse artist template", err)
		return
	}

	all, err := GetDetails(url + "/" + ID)
	if err != nil {
		RenderError(w, "Please try later", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, &all)
	if err != nil {
		RenderError(w, "Please try later", http.StatusInternalServerError)
		fmt.Println("Failed to execute artist template", err)
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		RenderError(w, "Please try later", http.StatusInternalServerError)
		fmt.Println("Failed to write in the buffer")
		return
	}
}
