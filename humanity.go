package humanity

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

// Number is an abstraction to support generic numbers as inputs
type Number interface {
	constraints.Integer | constraints.Float
}

// DurationFormat returns a human-readable string representation of nanosecond number
func DurationFormat[N Number](ns N) string {
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
func ByteFormat[N Number](numIn N) string {
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

// StringAsBytes will convert WHOLE NUMBER stringified byte-storage sizes to the corresponding int64 byte value.
// Decimals bomb by design.
// Suffixes above Petabyte are not supported due to int64 constraints.
// Suffixes must be ``[KMGTP]*?B)``.
// Case-insensitive.
func StringAsBytes(s string) (int64, error) {
	suffix := "B"
	units := []string{"", "K", "M", "G", "T"} // "P" caught  below
	re := regexp.MustCompile(`([\d.]+)([KMGTP]*?B)`)

	s = strings.ToUpper(s)
	parts := re.FindStringSubmatch(s)
	if lp := len(parts); lp == 0 {
		return -1, fmt.Errorf("no valid match for '%s', must be of the format '\\d[KMGTP]B'", s)
	} else if lp < 3 {
		return -1, fmt.Errorf("insufficient match for '%s', but this should be impossible", s)
	}

	n, err := strconv.Atoi(parts[1])
	if err != nil {
		return -1, err
	}

	num := int64(n)
	for _, unit := range units {
		if parts[2] == fmt.Sprintf("%s%s", unit, suffix) {
			return num, nil
		}
		num *= 1024
	}
	// Y
	return num, nil
}
