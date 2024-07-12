package main

import (
	"net/http"
)

func routes() *http.ServeMux {
	mux := http.NewServeMux()
	file_handler := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	http.Handle("/static/", file_handler)
	mux.HandleFunc("/", home)
	mux.HandleFunc("/ws", socket)
	return mux
}
