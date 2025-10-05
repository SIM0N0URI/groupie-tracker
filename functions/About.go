package functions

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func About(w http.ResponseWriter, r *http.Request) {
	ID := strings.TrimPrefix(r.URL.Path, "/artists/")

	if ID == "" {
		RenderError(w,"Page not found",http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(ID)
	if err != nil || id < 1 || id > 52 {
		RenderError(w,"Page not found",http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		RenderError(w,"Method not allowed",http.StatusMethodNotAllowed)
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

	err3 := tmpl.Execute(&buf, &all)
	if err3 != nil {
		RenderError(w, "Please try later", http.StatusInternalServerError)
		fmt.Println("Failed to execute artist template", err3)
		return
	}

	_, err4 := buf.WriteTo(w)
	if err4 != nil {
		RenderError(w, "Please try later", http.StatusInternalServerError)
		fmt.Println("Failed to write in the buffer")
		return
	}
}
