package main

import "fmt"

func main() {
	var totalAlbum, totalPossui int
    fmt.Scan(&totalAlbum)
	fmt.Scan(&totalPossui)

	mapSet := make(map[int]int)

    var repetidos []int 
    var faltando []int

	for i := 0; i < totalPossui; i++ {
		var value int
		fmt.Scan(&value)

		if _, ok := mapSet[value]; ok {
			repetidos = append(repetidos, value)
		}

		mapSet[value]++
	}

    for i := 1; i <= totalAlbum; i++ {
		_, ok := mapSet[i]
		if !ok {
			faltando = append(faltando, i)
		}
	}

    if len(repetidos) == 0 {
		fmt.Println("N")
	} else {
		for i, v := range repetidos {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(v)
		}
		fmt.Println()
	}

    if len(faltando) == 0 {
		fmt.Println("N")
	} else {
		for i, v := range faltando {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(v)
		}
		fmt.Println()
	}


}
