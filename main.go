package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/samir93bj/go-gorm-restapi/db"
	"github.com/samir93bj/go-gorm-restapi/routes"
)

func main() {

	db.DBconnection()

	router := mux.NewRouter()
	router.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":8080", router)
	println("Server started on port 8080")
}
