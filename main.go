package main

import (
	"log"
	"net/http"
)

func view(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello\n"))
}

func main() {
	http.HandleFunc("/view", view)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
