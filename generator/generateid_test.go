package generator

import (
	"testing"
	"fmt"

	//"github.com/hart87/goflake/components"
)

func TestGenerateTheIdentifier(t *testing.T) {
	result := GenerateIdentifier()
	fmt.Println(result)
	if result == "" {
		t.Errorf("Generated Identifier is blank.")
	}
}

func TestInt64ToStringForProperReturnType(t *testing.T) {
	result := ConvertInt64ToString(14591491198424)
	fmt.Println(result)
	if result != "14591491198424" {
		t.Errorf("Improper conversion type.")
	}
}

func TestIntToStringForProperReturnType(t *testing.T) {
	result := ConvertInt64ToString(14118)
	fmt.Println(result)
	if result != "14118" {
		t.Errorf("Improper conversion type.")
	}
}