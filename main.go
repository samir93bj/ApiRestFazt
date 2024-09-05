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
	db.DB.AutoMigrate(models.User{})

	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)

	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/users", routes.PostUserHandler).Methods("POST")

	http.ListenAndServe(":8080", router)
	println("Server started on port 8080")
}
