package functions

import (
	"fmt"
	"time"
)

func printDateInfo(datetime time.Time) {
	var (
		a int
	)

	fmt.Println("langauge")
	fmt.Println(datetime)
}

func PrintNow() {
	now := time.Now()
	printDateInfo(now)
}
