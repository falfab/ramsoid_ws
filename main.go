package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"ramsoid_ws/routes"
)

func main()  {
	router := mux.NewRouter()
	router.HandleFunc("/", routes.HomeHandler).Methods("GET")
	router.HandleFunc("/encryption", routes.EncryptionHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080",router))
}
