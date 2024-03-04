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

func TestJsonDoesntExist(t *testing.T) {
	id := GetId("adjnklsnjkladsnjk.json")

	if id != "" {
		t.Fatalf("The function allowed a JSON that didn't exist")
	}
}

func TestJsonIncorrectStructure(t *testing.T) {
	id := GetId("testwrongdata.json")

	if id != "" {
		t.Fatalf("The function allowed a JSON that was missing fields")
	}
}
