package generator

import (
	"strconv"
)

//ConvertInt64ToString convrts an int64 into a String
func ConvertInt64ToString(source int64) string {

	str := strconv.FormatInt(source, 10)
	return str
}

//ConvertIntToAString converts a int into a String
func ConvertIntToAString(source int) string {
	str := strconv.Itoa(source)
	return str
}
