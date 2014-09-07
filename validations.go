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

// func (scrivener *scrivener) AssertEqual(fieldName1 string, fieldName2 string) {
// scrivener.Assert(func(s interface{}) bool {
// value1 := scrivener.elem.FieldByName(fieldName1).String()
// value2 := scrivener.elem.FieldByName(fieldName2).String()

// return value1 == value2
// }, fieldName, "not_equal")
// }
