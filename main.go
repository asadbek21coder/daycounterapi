package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/asadbek21coder/daycounter/handler"
)

const (
	PORT = "7777"
)

func main() {
	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: nil,
	}

	fmt.Println("Server is running on port: ", PORT)
	http.Handle("/", handler.TokenMiddleware(http.HandlerFunc(handler.GetHomeHandler)))

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}

}
