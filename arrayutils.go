package go_utils

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func ScanArray2D(A, B int) [][]int {
	rows := make([][]int, A)
	for i := range rows {
		rows[i] = ScanArray(B)
	}
	return rows
}

func ScanArray(size int) []int {
	var tmp int
	array := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&tmp)
		array[i] = tmp
	}
	return array
}

//#endregion Array Input Helpers

func PrintArray(arr []int) {
	fmt.Printf("%v\n", arr)
}

func PrintArray2D(data [][]int) {
	for _, row := range data {
		fmt.Printf(" %v\n", row)
	}
}

func DupArray(array []int) []int {
	var buffer []int = make([]int, len(array))
	copy(buffer, array)
	return buffer
}

func DupStrArray(array []string) []string {
	var buffer []string = make([]string, len(array))
	copy(buffer, array)
	return buffer
}

func ArrayToString(arr []int) string {
	return strings.Join(strings.Fields(fmt.Sprint(arr)), ",")
}

func StringArrayToString(arr []string) string {
	return fmt.Sprintf("[%s]", JoinAndSurround(arr, ", ", "\""))
}

func RemoveArrayElement(arr []int, i int) []int {
	copy(arr[i:], arr[i+1:]) // Shift a[i+1:] left one index.
	// NOTE: data's still there, it's just this slice is smaller now
	return arr[:len(arr)-1] // Truncate slice.
}

func RemoveStringArrayElementIfExists(ss []string, s string) (bool, []string) {
	i := sort.SearchStrings(ss, s)
	if i < len(ss) && ss[i] == s {
		// doesn't exist?
		copy(ss[i:], ss[i+1:])      // Shift a[i+1:] left one index.
		return true, ss[:len(ss)-1] // Truncate slice.
	}

	return false, ss
}

func FindWithRegex(arr []string, regex_s string) []string {
	matches := make([]string, 0, 0)
	re := regexp.MustCompile(regex_s)
	for _, s := range arr {
		if re.FindString(s) != "" {
			matches = append(matches, s)
		}
	}

	return matches
}

func FindIndex(arr []string, searchStr string) int {
	for i, s := range arr {
		if searchStr == s {
			return i
		}
	}
	return -1
}

func FindIndexFast(ss []string, s string) int {
	i := sort.SearchStrings(ss, s)
	if i < len(ss) && ss[i] == s {
		return i
	}
	return -1
}

func ContainsString(arr []string, searchStr string) bool {
	for _, s := range arr {
		if searchStr == s {
			return true
		}
	}
	return false
}

func ContainsStringFast(ss []string, s string) bool {
	i := sort.SearchStrings(ss, s)
	return i < len(ss) && ss[i] == s
}

func ContainsSubtringFast(ss []string, s string) bool {
	i := sort.SearchStrings(ss, s)
	return i < len(ss) && strings.Contains(ss[i], s) // TODO: is this right?
}

func RemoveStringArrayElement(arr []string, i int) []string {
	copy(arr[i:], arr[i+1:]) // Shift a[i+1:] left one index.
	// NOTE: data's still there, it's just this slice is smaller now
	return arr[:len(arr)-1] // Truncate slice.
}

func InsertIntoSortedListIfNotThereAlready(ss []string, s string) (bool, []string) {
	i := sort.SearchStrings(ss, s)
	if i < len(ss) {
		// doesn't exist?
		if ss[i] == s {
			// already there, don't do anything
			return false, ss
		}
	}

	ss = append(ss, "")
	copy(ss[i+1:], ss[i:])

	// insert in order to keep sorted
	ss[i] = s
	return true, ss
}

func LastElement(arr []string) string {
	if len(arr) == 0 {
		return ""
	}
	return arr[len(arr)-1]
}

func RemoveEmpty(arr []string) []string {
	var rv []string
	for _, str := range arr {
		if str != "" {
			rv = append(rv, str)
		}
	}
	return rv
}
