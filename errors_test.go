package scrivener

import (
	"testing"
)

func TestErrorsAny(t *testing.T) {
	form := new(signup)
	scrivener := New(form)

	if scrivener.Errors.Any() {
		t.Error("Failed!")
	}
}
