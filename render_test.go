package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender(t *testing.T) {
	t.Parallel()

	type CustomResponse struct {
		Uno string `json:"uno"`
		Dos string `json:"DOS"`
	}

	expectedStatus := http.StatusBadGateway
	expectedBody := `{"uno":"one","DOS":"two"}`

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Render(w, expectedStatus, CustomResponse{
			Uno: "one",
			Dos: "two",
		})
	}).ServeHTTP(res, req)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	assertIntsEqual(t, expectedStatus, res.Code, "Render.Response.Code")
	assertStringsEqual(t, expectedBody, string(body), "Render.Response.Body")
}
