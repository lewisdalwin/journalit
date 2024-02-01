package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func routes() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/", home)
	router.HandlerFunc(http.MethodGet, "/about", about)
	router.HandlerFunc(http.MethodGet, "/entry/create", createEntry)
	return router
}
