package tools

import (
	"math/rand"
	"time"
)

var sourceUnixNano rand.Source

// UnixNano returns an int64 randomly choosen between 0 and r, using UnixNano as a Source
func RandUnixNano(r int64) int64 {
	if sourceUnixNano == nil {
		sourceUnixNano = rand.NewSource(time.Now().UnixNano())
	}
	return rand.New(sourceUnixNano).Int63n(r)
}

// DurationUnixNano same as UnixNano with time.Duration
func RandDurationUnixNano(r time.Duration) time.Duration {
	return time.Duration(RandUnixNano(int64(r)))
}
