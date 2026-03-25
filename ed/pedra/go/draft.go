package main

import "fmt"

func myJoin(lista []int, sep string) string {
	saida := ""

	for i, v := range lista {
		if i != 0 {
			saida += sep
		}
	}

	return saida
}

func main() {
	var num int
	fmt.Scan(&num)

	for i := 0; i < num; i++ {
		var n1, n2 int
		fmt.Scan(&n1, &n2)

	}
}
