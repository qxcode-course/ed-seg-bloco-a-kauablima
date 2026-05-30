package main

import "fmt"

type Queue struct {
	gasosa int
	dist   int
}

func Driven(vet []Queue) int {
	gasNow := 0
	start := 0

    for i := 0 ; i < len(vet); i++{
        gasNow += vet[i].gasosa - vet[i].dist

        if gasNow < 0 {
            start = i + 1 
            gasNow = 0 
        }
    }
    
	return start
}

func main() {
	var num int
	fmt.Scan(&num)

	vet := make([]Queue, num)

	for i := 0; i < num; i++ {
		fmt.Scan(&vet[i].gasosa, &vet[i].dist)
	}

	fmt.Println(Driven(vet))
}
