package server

import (
	"testing"
)

func TestRequest(t *testing.T) {
	word := PostId("12345")

	if word != `"Echo"` {
		t.Fatalf("The incorrect word was gotten")
	}
}

func TestBadId(t *testing.T) {
	word := PostId("abcde")

	if word != "" {
		t.Fatalf("The request was sent with the wrong type of ID")
	}
}
