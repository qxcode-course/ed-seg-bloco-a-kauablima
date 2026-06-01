package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	descasados := make(map[int]int)
	casais := 0

	for i := 0; i < num; i++ {
		var animal int
		fmt.Scan(&animal)

		par := -animal

		if descasados[par] > 0 {
			descasados[par]--
			casais++
		} else {
			descasados[animal]++
		}
	}

	fmt.Println(casais)
}
