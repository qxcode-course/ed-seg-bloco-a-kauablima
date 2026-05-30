package main

import (
	"fmt"
	"math"
)

type Players struct {
	p1 float64
	p2 float64
}

func Champion(vet []Players) int {
	winner := -1
	var count float64

	for i := 0; i < len(vet); i++ {
		if vet[i].p1 >= 10 && vet[i].p2 >= 10 {
			difference := math.Abs(vet[i].p1 - vet[i].p2)

			if winner == -1 || difference < count {
				count = difference
				winner = i
			}
		}
	}

	return winner
}

func main() {
	var num int
	fmt.Scan(&num)

	vet := make([]Players, num)
	for i := 0; i < num; i++ {
		fmt.Scan(&vet[i].p1, &vet[i].p2)
	}

	result := Champion(vet)
	if result != -1 {
		fmt.Println(Champion(vet))
	} else {
		fmt.Println("sem ganhador")
	}
}
