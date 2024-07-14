package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/start", startHandler)

	fmt.Printf("starting server at port: 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleNotFound(w http.ResponseWriter) {
	http.Error(w, "404 not found", http.StatusNotFound)
}

func validPathMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		validPaths := map[string]bool{
			"/":     true,
			"/form": true,
			"/start": true,
		}
		if !validPaths[r.URL.Path] {
			handleNotFound(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parseForm error: %v", err)
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name = %s\n", name)
	fmt.Fprintf(w, "address = %s\n", address)
}

func startHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusNotFound)
		return
	}
	fmt.Printf("initiated")
}