// Filename: cmd/web/main.go
package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// Configure server using command line arguments
	address := flag.String("addr", ":4000", "address of server")
	dsn := flag.String("db-dsn", "", "PostgreSQL DSN")
	flag.Parse()
	db, err := openDB(*dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Printf("database connection pool established")
	log.Printf("starting server on %s", *address)
	err = http.ListenAndServe(*address, routes())
	// We will only get here when the server stops (crashes)
	log.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
