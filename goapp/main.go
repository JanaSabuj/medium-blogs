package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print(r.URL.Path) // to log all incoming requests
	paramString := r.URL.Path[1:]
	if paramString == "" {
		paramString = "Everything!"
	}
	fmt.Fprintf(w, "Hi There !! I am running on port %s and I love %s!", os.Getenv("APPPORT"), paramString)
	// eg: hostname:port/ extracts Everything
	// eg: hostname:port/mangoes extracts mangoes
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APPPORT")), nil))
}
