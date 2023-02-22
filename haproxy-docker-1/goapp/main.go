package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi There !! I am running on port %s and I love %s!", os.Getenv("APPPORT"), r.URL.Path[1:]) // eg: hostname:port/mangoes extracts mangoes
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APPPORT")), nil))
}
