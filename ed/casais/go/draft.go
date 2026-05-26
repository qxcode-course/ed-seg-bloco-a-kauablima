package main

import "fmt"

func Converter(vet []int) {
	for i := range vet {
		if vet[i] < 0 {
			vet[i] = vet[i] * -1
		}
	}
}

func main() {
	var num int
	fmt.Scan(&num)

	vet := make([]int, num)
	for i := range num {
		var value int
		fmt.Scan(&value)

		vet[i] = value
	}

	Converter(vet)

	count := 0
	for i := 0; i < len(vet)-1; i++ {
		if vet[i] == vet[i+1] {
			count++
		}
	}

	fmt.Println(count)
}
