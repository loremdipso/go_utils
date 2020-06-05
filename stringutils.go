package go_utils

import (
	"fmt"
	"strings"
)

// Normal join includes the separator at the end, which is dumb
func Join(arr []string, delim string) string {
	var builder strings.Builder
	builder.Grow(len(arr) * (10 + len(delim))) // giving it a guess
	for index, value := range arr {
		if index < len(arr)-1 {
			fmt.Fprintf(&builder, "%s%s", value, delim)
		} else {
			fmt.Fprintf(&builder, "%s", value)
		}
	}
	return builder.String()
}

func JoinAndSurround(arr []string, delim string, surrounding string) string {
	var builder strings.Builder
	builder.Grow(len(arr) * (10 + len(delim))) // giving it a guess
	for index, value := range arr {
		if index < len(arr)-1 {
			fmt.Fprintf(&builder, "%s%s%s%s", surrounding, value, surrounding, delim)
		} else {
			fmt.Fprintf(&builder, "%s%s%s", surrounding, value, surrounding)
		}
	}
	return builder.String()
}

func FindLongest(arr []string) int {
	var maxLength int
	for _, value := range arr {
		length := len(value)
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}
