package humanity

import (
	"fmt"
)

// DurationFormat returns a human-readable string representation of nanosecond number
func DurationFormat(ns int64) string {
	suffix := "s"
	num := float64(ns)
	units := []string{"n", "u", "m"}
	for _, unit := range units {
		if num < 1000.0 {
			return fmt.Sprintf("%3.1f%s%s", num, unit, suffix)
		}
		num = (num / 1000)
	}
	return fmt.Sprintf("%.1f%s%s", num, "", suffix)
}

// ByteFormat returns a human-readable string representation of a byte count
func ByteFormat(numIn int64) string {
	suffix := "B"
	num := float64(numIn)
	units := []string{"", "K", "M", "G", "T", "P", "E", "Z"} // "Y" caught  below
	for _, unit := range units {
		if num < 1024.0 {
			return fmt.Sprintf("%3.1f%s%s", num, unit, suffix)
		}
		num = (num / 1024)
	}
	return fmt.Sprintf("%.1f%s%s", num, "Y", suffix)
}
