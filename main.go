package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.Handle("/iseven", isEven(http.HandlerFunc(getTime)))

	http.HandleFunc("/time", getTime)

	http.HandleFunc("/", func(w http.ResponseWriter, h *http.Request) {
		w.Write([]byte("hello world"))
	})

	log.Println("web server started at localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("unable to start web server", err)
	}
}

func getTime(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(time.Now().String()))
}

func isEven(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if time.Now().Second()%2 == 0 {
			h.ServeHTTP(w, r)
		}
		http.Error(w, "current time second is odd, returning HTTP error when odd. Refresh until is even", http.StatusInternalServerError)
	})
}
