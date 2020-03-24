package main

import (
	"fmt"
)

func main() {
	var speed int    // declare int var
	var heat float64 // declare float64 var
	var off bool     // declare bool var
	var brand string // declare string var

	var ( // other way of declaring multiple vars
		speed2, speed3 int
		heat2          float64
		off2           bool
		brand2         string
	)

	_ = speed // blank identifier

	fmt.Println(speed, speed2, speed3)
	fmt.Println(heat, heat2)
	fmt.Println(off, off2)
	fmt.Println(brand, brand2)

	fmt.Println(
		-50.5, -21., 34.2,
	)
	fmt.Println(
		true, false,
	)

	fmt.Println(
		speed, speed2,
	)

}
