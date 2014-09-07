package scrivener

import (
	"reflect"
)

type scrivener struct {
	elem   reflect.Value
	Errors []string
}

func New(s interface{}) *scrivener {
	form := new(scrivener)

	form.Errors = make([]string, 0)
	form.elem = reflect.ValueOf(s).Elem()

	return form
}
