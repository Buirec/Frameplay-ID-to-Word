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
