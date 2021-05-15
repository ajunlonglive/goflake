package generator

import (
	"github.com/hart87/goflake/components"
)

//GenerateIdentifier is the function that takes all of the components and generates the
//unique identifier requested. Sample : 0-14591498880493-0311-425471-43Que
func GenerateIdentifier() string {

	//1 or 0 for auditing purposes. 0 Default
	checkDigit := components.AddCheckDigit(false)

	//Milliseconds ellapsed since your epoch. Set in parameter below
	//1609459201000 = 01/01/2021 00:001 hrs GMT. milliseconds
	epochInt := components.TimeSinceOrganizationEpoch(1609459201000)
	epochStr := ConvertInt64ToString(epochInt)

	//A naming/numbering convention for the service.
	serviceNum := components.SetServiceNumber("0311")

	//Process ID for the machine this runs on
	pidInt := components.ObtainPID()
	pidStr := ConvertIntToAString(pidInt)

	//equence number
	seq := components.RandString()

	//Concatenate the pieces to form a Identifier.
	result := checkDigit + "-" + epochStr + "-" + serviceNum + "-" + pidStr + "-" + seq
	return result
}
