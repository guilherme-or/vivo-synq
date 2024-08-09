package handler

import (
	"encoding/json"
	"net/http"
)

type JSON map[string]interface{}

func setJSONHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func writeJSON(w http.ResponseWriter, code int, data any) {
	setJSONHeader(w)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func writeErr(w http.ResponseWriter, err *AppErr) {
	writeJSON(w, err.Status, err)
}
