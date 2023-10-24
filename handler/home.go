package handler

import (
	"html/template"
	"net/http"
	"time"
)

type TemplateData struct {
	DaysRemaining int
}

func GetHomeHandler(w http.ResponseWriter, r *http.Request) {
	target := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	remain := int(target.Sub(time.Now()).Hours() / 24)

	data := TemplateData{DaysRemaining: remain}

	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
