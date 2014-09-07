package scrivener

func (scrivener *scrivener) Assert(fn func(interface{}) bool, key string, message string) bool {
	pass := fn(scrivener.s)
	if !pass {
		scrivener.Errors[key] = append(scrivener.Errors[key], message)
		return false
	}
	return true
}

func (scrivener *scrivener) AssertPresent(fieldName string) bool {
	return scrivener.Assert(func(_ interface{}) bool {
		value := scrivener.elem.FieldByName(fieldName).String()

		return len(value) > 0
	}, fieldName, "not_present")
}

func (scrivener *scrivener) AssertEqualString(fieldName string, expected string) bool {
	return scrivener.Assert(func(_ interface{}) bool {
		actual := scrivener.elem.FieldByName(fieldName).String()

		return actual == expected
	}, fieldName, "not_equal")
}
