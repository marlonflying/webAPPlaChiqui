package main

import (
	"log"
	"net/http"
)

//home handler function with a string as the response
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home response from La Chiqui webAPP"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe("4000", mux)
	log.Fatal(err)
}
