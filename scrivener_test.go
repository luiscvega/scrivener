package scrivener

import (
	"reflect"
	"testing"
)

type signup struct {
	Email           string `name:"email"`
	FirstName       string `name:"first_name"`
	LastName        string `name:"last_name"`
	Password        string `name:"password"`
	ConfirmPassword string `name:"confirm_password"`
}

func TestAssertPresent(t *testing.T) {
	form := new(signup)
	scrivener := New(form)

	scrivener.AssertPresent("FirstName")
	if !reflect.DeepEqual(scrivener.Errors, map[string][]string{"FirstName": {"not_present"}}) {
		t.Error("Failed!")
	}
}

func TestAssertEqualString(t *testing.T) {
	form := &signup{LastName: "Vega"}
	scrivener := New(form)

	scrivener.AssertEqualString("LastName", "Cancio")
	if !reflect.DeepEqual(scrivener.Errors, map[string][]string{"LastName": {"not_equal"}}) {
		t.Error("Failed!")
	}
}
