package routes

import (
	"net/http"
	"encoding/json"
	"ramsoid_ws/models"
)

func EncryptionHandler(w http.ResponseWriter, r *http.Request) {
	println("----- New Encrypt request ------")
	decoder := json.NewDecoder(r.Body)
	var record models.Record
	err := decoder.Decode(&record)
	if err != nil {
		println("ERROR: " + err.Error())
	}
	println("ID: " + record.ID)
	println("KEY: " + record.Key)
	println("--------------------------------")
	w.WriteHeader(200)
}

func DecryptionHandler(w http.ResponseWriter, r *http.Request) {
	println("----- New Decrypt request ------")
	decoder := json.NewDecoder(r.Body)
	var record models.Record
	err := decoder.Decode(&record)
	if err != nil {
		println("ERROR: " + err.Error())
	}
	println("ID: " + record.ID)
	println("KEY: " + record.Key)

}