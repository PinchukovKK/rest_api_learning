package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var task string

type RequstBody struct {
	Massage string `json:"massage"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, ", task)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var req RequstBody
	json.NewDecoder(r.Body).Decode(&req)
	task = req.Massage
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api", PostHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
