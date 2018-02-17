package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"ramsoid_ws/routes"
	"ramsoid_ws/db"
)

func main()  {

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/", routes.HomeHandler).Methods("GET")
	router.HandleFunc("/enc", routes.EncryptionHandler).Methods("POST")
	router.HandleFunc("/dec", routes.DecryptionHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080",router))
}
