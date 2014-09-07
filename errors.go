package scrivener

type Errors map[string][]string

func (e Errors) Any() bool {
	if len(e) > 0 {
		return true
	}
	return false
}

