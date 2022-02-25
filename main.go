package main

import (
	"api-golang/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/students", controllers.GetStudents).Methods("GET")
	// router.HandleFunc("/students/{id}", GetPerson).Methods("GET")
	// router.HandleFunc("/students/{id}", CreatePerson).Methods("POST")
	// router.HandleFunc("/students/{id}", DeletePerson).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}
