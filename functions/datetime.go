package functions

import (
	"fmt"
	"time"
)

func printDateInfo(datetime time.Time) {
	fmt.Println(datetime)
}

// PrintNow prints current time
func PrintNow() {
	now := time.Now()
	printDateInfo(now)
}
