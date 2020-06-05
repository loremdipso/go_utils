package go_utils

import (
	"fmt"
	"time"
)

//#region Timer
func Timer(fn func(), description string, iterations int) {
	start := time.Now()
	for i := 0; i < iterations; i++ {
		fn()
	}
	fmt.Printf("%s\n%v\n\n", description, time.Since(start))
}

//#endregion Timer
