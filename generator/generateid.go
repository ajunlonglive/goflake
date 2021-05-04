package generator

import (

	"github.com/hart87/goflake/components"
)

//Sample : 0-14591498880493-0311-425471-43Que
func GenerateIdentifier() string {

	checkDigit := components.AddCheckDigit(false) 
	
	epochInt := components.TimeSinceOrganizationEpoch()
	epochStr := ConvertInt64ToString(epochInt)

	serviceNum := components.SetServiceNumber("0311") // users, REST, AWS

	pidInt := components.ObtainPID()
	pidStr := ConvertIntToAString(pidInt)

	seq := components.RandString()

	//Concatenate the pieces to form a Identifier.
	result := checkDigit + "-" + epochStr + "-" + serviceNum + "-" + pidStr + "-" + seq
	return result
}

