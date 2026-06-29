package auth

import (
  "net/http"
	"encoding/json"
)

type AuthPayload struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func Login (w http.ResponseWriter, r *http.Request) {
	var loginPayload AuthPayload

	if (json.NewDecoder(r.Body).Decode(&loginPayload) != nil){
		http.Error(w, "Invalid JSON Body", http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}

func Register (w http.ResponseWriter, r *http.Request){
	var registerPayload AuthPayload

	if (json.NewDecoder(r.Body).Decode(&registerPayload) != nil){
		http.Error(w, "Invalid JSON Body", http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
}


