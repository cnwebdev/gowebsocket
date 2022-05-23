package main

import (
	"net/http"
	"websocket/ws/internal/handlers"

	"github.com/gorilla/mux"
)

func routes() http.Handler {
	mux := mux.NewRouter()

	fileServer := http.FileServer(http.Dir("assets"))
	mux.PathPrefix("/assets").Handler(http.StripPrefix("/assets", fileServer))
	// http.Handle("/html/css/*", http.StripPrefix("/html/css", fileServer))

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("ws", handlers.WsEndpoint)

	return mux
}
