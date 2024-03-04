package server

import (
	"testing"
)

func TestCorrectJson(t *testing.T) {
	id := GetId("test.json")

	if id != "12345" {
		t.Fatalf("The id is wrong")
	}
}

func TestIncorrectJson(t *testing.T) {
	id := GetId("testfail.json")

	if id != "" {
		t.Fatalf("The function allowed an invalid id")
	}
}
