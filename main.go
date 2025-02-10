package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"main.go/internal/database"
	"main.go/internal/handlers"
	"main.go/internal/taskService"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewTaskService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api", handler.GetTaskHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
