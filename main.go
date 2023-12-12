package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}
	w.Write([]byte("Hello fom SnippetBox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 0 {
		http.NotFound(w,r)
		return
	}
	// w.Write([]byte("Display a specific snippet with ID here....."))
	fmt.Fprintf(w, "Display a specific snippet with ID %d .....", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create snippets here...."))
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/create", snippetCreate )
	mux.HandleFunc("/snippet/view", snippetView)

	log.Print("starting server at localhost:4000")


	err := http.ListenAndServe("localhost:4000", mux)
	log.Fatal(err)
}