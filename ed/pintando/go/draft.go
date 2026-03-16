package main

import (
	"fmt"
	"math"
)

func formulaHeron(a float64, b float64, c float64) float64 {
	p := (a + b + c) / 2
	value := p * (p - a) * (p - b) * (p - c)
	area := math.Round(math.Sqrt(value)*100) / 100

	return area
}

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	fmt.Printf("%.2f\n", formulaHeron(a, b, c))
}
