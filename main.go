package main

import (
	"fmt"

	hmmm "github.com/ramamimu/golang-modules/hello"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(hmmm.GetHello("ramamimu"))
	hmmm.JustPrint()
}
