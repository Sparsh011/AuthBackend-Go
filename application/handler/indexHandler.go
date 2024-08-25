package handler

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Response struct {
	Message string `json:"message"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response := Response{
		Message: "Welcome to the index page",
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
