package main

import (
	"math/rand"
	"net/http"
	"time"
)

// ErrorResponse defines the response of a generic error.
type ErrorResponse struct {
	Error string `json:"error"`
}

// RichStatusResponse defines the response of the RichStatusHandfler.
type RichStatusResponse struct {
	Hostname string    `json:"hostname"`
	OK       bool      `json:"ok"`
	Quote    string    `json:"quote"`
	Time     time.Time `json:"time"`
	Version  string    `json:"version"`
}

// RichStatusHandler provides a JSON response mimicing values one might want
// from a status or equivalent endpoint. It also provides a fun quote.
func RichStatusHandler(hostname string, quotes []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			ErrorHandler(http.StatusNotFound)(w, r)
			return
		}

		Render(w, http.StatusOK, RichStatusResponse{
			Hostname: hostname,
			OK:       true,
			Quote:    quotes[rand.Intn(len(quotes))],
			Time:     time.Now(),
			Version:  "2.0",
		})
	}
}

// ErrorHandler provides a helper for responding with a generic http error.
func ErrorHandler(code int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Render(w, code, ErrorResponse{
			Error: http.StatusText(code),
		})
	}
}

// FileHandler provides a helper for serving a file.
func FileHandler(f string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, f)
	}
}
