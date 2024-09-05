package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/samir93bj/go-gorm-restapi/db"
	"github.com/samir93bj/go-gorm-restapi/models"
	"github.com/samir93bj/go-gorm-restapi/routes"
)

func main() {

	db.DBconnection()
	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Task{})

	router := mux.NewRouter()

	s := router.PathPrefix("/api").Subrouter()

	s.HandleFunc("/", routes.HomeHandler)

	// tasks routes
	s.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	s.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	s.HandleFunc("/tasks", routes.PostTaskHandler).Methods("POST")

	s.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	s.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	s.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")
	s.HandleFunc("/users", routes.PostUserHandler).Methods("POST")

	http.ListenAndServe(":8080", router)
	println("Server started on port 8080")
}
