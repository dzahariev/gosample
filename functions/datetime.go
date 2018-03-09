package functions

import (
	"fmt"
	"time"
)

func printDateInfo(datetime time.Time) {
	fmt.Println(datetime)
}

func PrintNow() {
	now := time.Now()
	printDateInfo(now)
}
