package main

import "fmt"

func handlePrint(vetor []int, swordIdx int) {
	fmt.Print("[")
	for i, val := range vetor {
		if i == swordIdx {
			fmt.Printf(" %v>", val)
		} else {
			fmt.Printf(" %v", val)
		}
	}
	fmt.Println(" ]")
}

func main() {
	var num, player int
	fmt.Scan(&num, &player)

	vetor := make([]int, num)
	for i := 0; i < num; i++ {
		vetor[i] = i + 1
	}

	swordIdx := 0
	for i, v := range vetor {
		if v == player {
			swordIdx = i
			break
		}
	}

	for len(vetor) > 1 {
		handlePrint(vetor, swordIdx)

		killIdx := (swordIdx + 1) % len(vetor)

		vetor = append(vetor[:killIdx], vetor[killIdx+1:]...)

		swordIdx = killIdx % len(vetor)
	}
	handlePrint(vetor, swordIdx)
}
