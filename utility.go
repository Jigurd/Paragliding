package main

import (
	"fmt"
	"github.com/marni/goigc"
	"regexp"
	"time"
)

//isNumeric checks whether a given string is numeric
func isNumeric(s string) bool {
	value, _ := regexp.MatchString("[0-9]+", s)
	return value
}

//Min returns the lower of the two int arguments
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//Min64 is similar to Min, but for Int64s.
func Min64(a, b int64) int64 {
	if a < b {
		return int64(a)
	}
	return int64(b)
}

//FindIndex returns index of the id as int and whether things succeeded as bool
func FindIndex(slice []Track, id int64) (int, bool) {
	AllOK := true
	for i := 0; i < len(slice); i++ {
		if slice[i].Timestamp == id {
			return i, AllOK
		}
	}
	AllOK = false
	return -1, AllOK
}

//IsInSlice looks for a specific id in a slice of tracks
func IsInSlice(slice []Track, id int64) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i].Timestamp == id {
			return true //if the id is found, return true
		}
	}
	return false //if it is not found in the array, return false
}

//Millisec returns monotonic time as an int64
func Millisec() int64 {
	return time.Now().UnixNano() / 1000000 //
}

//TotalDistance Finds the total distance of an IGC track
func TotalDistance(t igc.Track) string {
	track := t
	totalDistance := 0.0
	for i := 0; i < len(track.Points)-1; i++ {
		totalDistance += track.Points[i].Distance(track.Points[i+1])
	}

	return fmt.Sprintf("%f", totalDistance)
}
