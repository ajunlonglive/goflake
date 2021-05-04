package generator

import (
	"strconv"
)

func ConvertInt64ToString(source int64) string {

	str := strconv.FormatInt(source, 10)
	return str
}

func ConvertIntToAString(source int) string {
	str := strconv.Itoa(source)
	return str
}