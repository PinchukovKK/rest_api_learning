package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tasks []Task
	if err := DB.Find(&tasks).Error; err != nil {
		http.Error(w, "Failed to get task", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var tasks Task
	json.NewDecoder(r.Body).Decode(&tasks)

	if err := DB.Create(&tasks).Error; err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]

	var updatedTask Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var task Task
	if err := DB.First(&task, taskID).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	if updatedTask.Task != "" {
		task.Task = updatedTask.Task
	}
	task.IsDone = updatedTask.IsDone

	if err := DB.Save(&task).Error; err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(task)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]

	var task Task
	if err := DB.First(&task, taskID).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	if err := DB.Delete(&task).Error; err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Task deleted successfully")
}

func main() {
	InitDB()

	if err := DB.AutoMigrate(&Task{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/api", GetHandler).Methods("GET")
	router.HandleFunc("/api/post", PostHandler).Methods("POST")
	router.HandleFunc("/api/task/{id}", PatchHandler).Methods("PATCH")
	router.HandleFunc("/api/delete/{id}", DeleteHandler).Methods("DELETE")

	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
