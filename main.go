package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/junxxx/go-playground/response"
)

type User struct {
	NickName string
}

// Get
func view(w http.ResponseWriter, r *http.Request) {
	b, _ := response.Json(0, map[string]int{"age": 29, "exp": 69}, "")
	w.Write(b)
}

// POST create a new user
func create(w http.ResponseWriter, r *http.Request) {
	if method := r.Method; method != http.MethodPost {
		b, _ := response.Json(1, "", "Method not allowed!")
		w.Write(b)
		return
	}
	defer r.Body.Close()
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}
	b, _ := response.Json(0, user, "User created!")
	w.Write(b)
	log.Println(user)
}

func main() {
	http.HandleFunc("/view", view)
	http.HandleFunc("/create", create)

	log.Fatal(http.ListenAndServe(":8888", nil))
}
