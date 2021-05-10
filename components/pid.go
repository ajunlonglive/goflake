package components

import (
	"os"
)

//Obtains the process ID of the machine running an operation for the identifier.
func ObtainPID() int {
	result := os.Getpid()
	return result
}
