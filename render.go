package main

import (
	"encoding/json"
	"net/http"
)

// Render provides a convenient JSON rendering method that ensures written
// response values are marshaled and written, otherwise triggers an error.
func Render(w http.ResponseWriter, code int, v interface{}) {
	bytes, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if _, err := w.Write(bytes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
