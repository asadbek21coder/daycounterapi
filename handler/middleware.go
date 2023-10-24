package handler

import (
	"html/template"
	"net/http"
)

const authToken = "hello" // random uuid

func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token != authToken {
			tmpl, err := template.ParseFiles("templates/unauthorized.html")
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, nil)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}

			return
		}

		next.ServeHTTP(w, r)
	})
}
