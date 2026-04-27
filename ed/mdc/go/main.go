package main

import (
	"fmt"
)

func mdc(a, b int) int {
	r := a % b

	mdc(a, b)
	return r
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(mdc(a, b))
}
