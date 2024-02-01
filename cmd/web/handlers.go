package main

import (
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welome to journalit!"))
}
func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the about page..."))
}

func createEntry(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Add a new journal entry..."))
}
