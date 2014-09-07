package scrivener

import (
	"reflect"
)

type scrivener struct {
	s      interface{}
	elem   reflect.Value
	Errors map[string][]string
}

func New(s interface{}) *scrivener {
	form := new(scrivener)

	form.s = s
	form.elem = reflect.ValueOf(s).Elem()

	form.Errors = make(map[string][]string, 0)

	return form
}
