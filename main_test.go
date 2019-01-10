package main

import (
	"fmt"
	"os"
	"testing"
)

func assertIntsEqual(t *testing.T, expected int, actual int, name string) {
	if expected != actual {
		t.Error(fmt.Sprintf("%s: got %d, want %d", name, actual, expected))
	}
}

func assertStringsEqual(t *testing.T, expected string, actual string, name string) {
	if expected != actual {
		t.Error(fmt.Sprintf("%s: got %s, want %s", name, actual, expected))
	}
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}
