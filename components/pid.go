package components

import(
	"os"
)

func ObtainPID() int {
	result := os.Getpid()
	return result
}
