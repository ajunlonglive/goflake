package components

import (
	"time"
)

func TimeSinceOrganizationEpoch() int64 {
	var orgEpoch int64 = 1609459201000 //01/01/2021 00:001 hrs GMT. milliseconds
	
	now := time.Now()
	nanos := now.UnixNano()
	millis := nanos / 100000

	timeSince := millis - orgEpoch
	return timeSince
}