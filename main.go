package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/junxxx/go-playground/response"
)

type User struct {
	NickName string
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Get
func view(w http.ResponseWriter, r *http.Request) {
	b, _ := response.Json(0, map[string]int{"age": 29, "exp": 69}, "")
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

// POST create a new user
func create(w http.ResponseWriter, r *http.Request) {
	if method := r.Method; method != http.MethodPost {
		b, _ := response.Json(1, "", "Method not allowed!")
		w.Header().Set("Content-Type", "application/json")
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
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
	log.Println(user)
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	for {
		// read a new message
		msgType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(message))
		if err := conn.WriteMessage(msgType, message); err != nil {
			log.Println(err)
			return
		}
	}

}

func main() {
	http.HandleFunc("/view", view)
	http.HandleFunc("/create", create)
	http.HandleFunc("/ws", serveWs)

	log.Fatal(http.ListenAndServe(":8888", nil))
}
