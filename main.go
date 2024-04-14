package main

import (
	"log"
	"net/http"

	"github.com/junxxx/go-playground/response"
)

func view(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello\n"))
	b, _ := response.Json(map[string]int{"age": 29, "exp": 69})
	w.Write(b)
}

func main() {
	http.HandleFunc("/view", view)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
