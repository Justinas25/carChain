package main

import (
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "http://www.google.com", 301)
}

func main() {
	http.HandleFunc("/", redirect)
}
