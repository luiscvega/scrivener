package scrivener

func (scrivener *scrivener) Assert(fn func(interface{}) bool, key string, message string) {
	result := fn(scrivener.s)
	if !result {
		scrivener.Errors[key] = append(scrivener.Errors[key], message)
	}
}

func (scrivener *scrivener) AssertPresent(fieldName string) {
	scrivener.Assert(func(s interface{}) bool {
		value := scrivener.elem.FieldByName(fieldName).String()

		return len(value) > 0
	}, fieldName, "not_present")
}

func (scrivener *scrivener) AssertEqualString(fieldName string, expected string) {
	scrivener.Assert(func(s interface{}) bool {
		actual := scrivener.elem.FieldByName(fieldName).String()

		return actual == expected
	}, fieldName, "not_equal")
}
