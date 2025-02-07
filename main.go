package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	if err := DB.Find(&tasks).Error; err != nil {
		http.Error(w, "Failed to create task1", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var tasks Task
	json.NewDecoder(r.Body).Decode(&tasks)

	if err := DB.Create(&tasks).Error; err != nil {
		http.Error(w, "Failed to create task2", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func main() {
	InitDB()

	if err := DB.AutoMigrate(&Task{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/api", GetHandler).Methods("GET")
	router.HandleFunc("/api/post", PostHandler).Methods("POST")
	
	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
