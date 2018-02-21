package routes

import (
	"net/http"
	"encoding/json"
	"ramsoid_ws/models"
	"ramsoid_ws/db"
)

func EncryptionHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var record models.Record
	err := decoder.Decode(&record)
	if err != nil {
		println("ERROR: " + err.Error())
	}
	println("\nPOST /enc:\n\t Key: " + record.Key + "\n\tId: " + record.ID)

	db.InsertRecord(&record)

	println("SUCCESS! new record inserted")

	println("=======================================================")

	w.WriteHeader(200)
}

func DecryptionHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var record models.Record
	err := decoder.Decode(&record)
	if err != nil {
		println("ERROR: " + err.Error())
	}
	println("\nPOST /dec:\n\tKey: " + record.Key + "\n\tId: " + record.ID)

	defer println("=======================================================")

	if db.CheckKey(&record) {
		println("SUCCESS! key and id matches")
		w.WriteHeader(200)
	} else {
		println("FAIL! key and id DO NOT matches")
		w.WriteHeader(999)
	}

}