// Filename: cmd/web/main.go
package main

import (
	"flag"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welome to journalit!"))
}

func main() {
	// Configure server using command line arguments
	address := flag.String("addr", ":4000", "address of server")
	flag.Parse()
	router := http.NewServeMux()
	router.HandleFunc("/", home)
	log.Printf("starting server on %s", *address)
	err := http.ListenAndServe(*address, router)
	// We will only get here when the server stops (crashes)
	log.Fatal(err)
}
