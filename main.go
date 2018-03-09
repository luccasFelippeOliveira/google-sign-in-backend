package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Message message model
type Message struct {
	Msg string
}

type User struct {
	Email  string
	Name   string
	UUID   int
	GToken string
}

var users []User

// Initialize user list.
func init() {
	users = make([]User, 2)
}

func login(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Auth-Token")
	if token != "" {
		// Validate token using google api.
	}
}

// TODO: Generate JWT
func generateJWT() {

}

// A request so pode ser feita se mandar a autorização.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	handleCors(&w)
	myMsg := Message{
		Msg: "Hello World",
	}

	msg, err := json.Marshal(myMsg)
	if err != nil {
		fmt.Println("Error encoding Msg")
	}
	fmt.Println("Encoding msg = " + myMsg.Msg)
	w.Write([]byte(msg))
}

// Check if token supplied by user is valid.

func handleCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "X-Auth-Code")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Starting server at 0.0.0.:3000")
	http.ListenAndServe(":3000", nil)
}
