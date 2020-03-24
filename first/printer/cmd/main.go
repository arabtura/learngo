package main

import (
	"fmt"
	"github/arabtura/learngo/first/printer"
	"github/arabtura/learngo/first/version"
)

//import "github/arabtura/learngo/first/printer"

func main() {
	printer.Hello()
	fmt.Println(version.Version())
}
