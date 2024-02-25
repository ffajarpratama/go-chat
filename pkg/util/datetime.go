package util

import (
	"time"
)

var Timezone *time.Location

func SetTimeZone(location string) {
	tz, err := time.LoadLocation(location)
	if err != nil {
		panic(err)
	}

	Timezone = tz
}
