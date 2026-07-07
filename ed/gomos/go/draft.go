package main

import "fmt"

type Pos struct {
	x int
	y int
}

func main() {
	var n int
	fmt.Scan(&n)
	var w string
	fmt.Scan(&w)

	vet := make([]Pos, n)
    
	for i := 0; i < n; i++ {
		fmt.Scan(&vet[i].x, &vet[i].y)
	}

	var x, y int
	fmt.Scan(&x, &y)

	for i :=n-1; i > 0; i-- {
		vet[i] = vet[i-1]
	}

	switch w {
	case "L":
		vet[0].x--
	case "R":
		vet[0].x++
	case "U":
		vet[0].y--
	case "D":
		vet[0].y++
	}

	for i := 0; i < n; i++ {
		fmt.Println(vet[i].x, vet[i].y)
	}

}
