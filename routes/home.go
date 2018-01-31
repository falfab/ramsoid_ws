package routes

import (
	"net/http"
	"encoding/json"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Server is running!")
}
