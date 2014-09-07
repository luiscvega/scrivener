package scrivener

func (scrivener *scrivener) AssertPresent(fieldName string) {
	result := scrivener.elem.FieldByName(fieldName).String()

	if result == "" {
		scrivener.Errors = append(scrivener.Errors, fieldName+" can't be blank!")
	}
}

func (scrivener *scrivener) AssertEqual(fieldName1 string, fieldName2 string) {
	result1 := scrivener.elem.FieldByName(fieldName1).String()
	result2 := scrivener.elem.FieldByName(fieldName2).String()

	if result1 != result2 {
		scrivener.Errors = append(scrivener.Errors, fieldName1+" and "+fieldName2+" do not match!")
	}
}
