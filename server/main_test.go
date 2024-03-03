package main

import (
	"testing"
)

func TestCorrectJson(t *testing.T) {
	id := getId("test.json")

	if id != "12345" {
		t.Fatalf("The id is wrong")
	}
}

func TestIncorrectJson(t *testing.T) {
	id := getId("testfail.json")

	if id != "false" {
		t.Fatalf("The function allowed an invalid id")
	}
}

func TestRequest(t *testing.T) {
	word := postId("12345")

	if word != `"Echo"` {
		t.Fatalf("The incorrect word was gotten")
	}
}
