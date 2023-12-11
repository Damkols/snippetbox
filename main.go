package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request ) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from the Homepage"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display snippets here...."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}
	w.Write([]byte("Creating snippets here....."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/snippet/view", snippetView)

	log.Println("Server running at localhost:4000")

	err := http.ListenAndServe("localhost:4000", mux)
	log.Fatal(err)
}