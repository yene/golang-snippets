// You can import this package into current package without namespace with the .
// import . "misc_utils"

package misc_utils

import (
	"fmt"
	"time"
)

// tracking the execution time of a function
// by running: defer TrackTime(time.Now())
func TrackTime(pre time.Time) time.Duration {
	elapsed := time.Since(pre)
	fmt.Println("elapsed:", elapsed)

	return elapsed
}
