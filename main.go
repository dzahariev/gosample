package main

import (
	"fmt"
	"time"
	"github.com/dzahariev/gosample/functions"
)

func init() {
	fmt.Println("Initialization ... %d")
}

func Dummy(dummy string) string {
	return dummy
}

func main() {
	functions.PrintNow()
	fmt.Println("5 + 11 =", functions.Plus(5, 11))
	fmt.Println("11 - 5 =", functions.Minus(11, 5))
}
