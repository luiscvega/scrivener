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

var form = new(signup)

func TestAssert(t *testing.T) {
	form.Email = "luiscvega@gmail.com"
	form.Password = "pass1234"
	form.ConfirmPassword = "pass1235"

	scrivener := New(form)

	scrivener.AssertPresent("FirstName")

	if !reflect.DeepEqual(scrivener.Errors, map[string][]string{"FirstName": {"not_present"}}) {
		t.Error("Failed!")
	}
}
