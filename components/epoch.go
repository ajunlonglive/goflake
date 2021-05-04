package components

import (
	"time"
)

func TimeSinceOrganizationEpoch(orgEpoch int64) int64 {
	now := time.Now()
	nanos := now.UnixNano()
	millis := nanos / 100000

	timeSince := millis - orgEpoch
	return timeSince
}