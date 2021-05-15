package components

import ()

//AddCheckDigit creates a check digit for an identifier
func AddCheckDigit(condition bool) string {
	if condition != true {
		return "0"
	}
	return "1"
}
