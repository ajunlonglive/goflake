package components

import (
	"testing"
	"fmt"
)

func TestObtainPID(t *testing.T) {

	result := ObtainPID()
	fmt.Println(result)
	if result == 0 {
		t.Errorf("No Process ID obtained.")
	}
}

func TestTimeSinceOrganizationEpoch(t *testing.T) {

	result := TimeSinceOrganizationEpoch()
	fmt.Println(result)
	if result <= 1000000  {
		t.Errorf("Improper Epoch Time calculated.")
	}
}

func TestCheckDigitBooleanOperationFalse(t *testing.T) {
	result := AddCheckDigit(false)
	fmt.Println(result)
	if result !=  "0" {
		t.Errorf("Boolean false operation did not set properly")
	}
}

func TestCheckDigitBooleanOperationTrue(t *testing.T) {
	result := AddCheckDigit(true)
	fmt.Println(result)
	if result !=  "1" {
		t.Errorf("Boolean true operation did not set properly")
	}
}

func TestSetServiceNumberNotNil(t *testing.T) {
	result := SetServiceNumber("0311")
	fmt.Println(result)
	if result != "0311" {
		t.Errorf("Set Service Number not creating string properly")
	}
}

func TestRandStringCreatesRandomString(t *testing.T) {
	result := RandString()
	fmt.Println(result)
	if result == "" {
		t.Errorf("No random string constructed or returned.")
	}
}
