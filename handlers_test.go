package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRichStatusHandler(t *testing.T) {
	t.Parallel()

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	RichStatusHandler("mock-host", []string{"mock-quote"}).ServeHTTP(res, req)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var v RichStatusResponse
	if err := json.Unmarshal(body, &v); err != nil {
		t.Fatal(err)
	}

	assertIntsEqual(t, http.StatusOK, res.Code, "RichStatusHandler.Response.Code")
	assertStringsEqual(t, "mock-host", v.Hostname, "RichStatusHandler.Response.Body.Hostname")
	assertStringsEqual(t, "mock-quote", v.Quote, "RichStatusHandler.Response.Body.Quote")
}

func TestRichStatusHandlerNotFound(t *testing.T) {
	t.Parallel()

	req, err := http.NewRequest(http.MethodGet, "/path", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	RichStatusHandler("", []string{}).ServeHTTP(res, req)

	assertIntsEqual(t, http.StatusNotFound, res.Code, "TestRichStatusHandlerNotFound.Response.Code")
}

func TestErrorHandler(t *testing.T) {
	t.Parallel()

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	ErrorHandler(http.StatusNotFound).ServeHTTP(res, req)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var v ErrorResponse
	if err := json.Unmarshal(body, &v); err != nil {
		t.Fatal(err)
	}

	assertIntsEqual(t, http.StatusNotFound, res.Code, "ErrorHandler.Response.Code")
	assertStringsEqual(t, http.StatusText(http.StatusNotFound), v.Error, "ErrorHandler.Response.Body.Error")
}

func TestFileHandler(t *testing.T) {
	t.Parallel()

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	FileHandler("./main_test.go").ServeHTTP(res, req)

	assertIntsEqual(t, http.StatusOK, res.Code, "FileHandler.Response.Code")
}
