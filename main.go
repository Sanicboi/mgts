package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		message := r.URL.Query().Get("message")
		fmt.Fprintln(w, "Hello, "+message)
	})

	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
