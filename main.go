package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler)

	http.ListenAndServe(":8080", router)
	println("Server started on port 8080")
}
