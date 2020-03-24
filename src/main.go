package main

import (
	"log"
	"net/http"
)

func main() {

	myhttp := http.NewServeMux()
	fs := http.FileServer(http.Dir("./views/"))
	myhttp.Handle("/", http.StripPrefix("", fs))

	myhttp.HandleFunc("/socket", socketReaderCreate)

	log.Println("http://localhost:8080")
	http.ListenAndServe(":8080", myhttp)
}
