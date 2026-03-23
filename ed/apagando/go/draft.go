package main

import (
	"fmt"
	"sort"
)

func main() {
	var people int
	fmt.Scan(&people)

	queue := make(map[int]int)

	for i := 0; i < people; i++ {
		var value int
		fmt.Scan(&value)
		queue[i] = value
	}

	var leaveQueue int
	fmt.Scan(&leaveQueue)

	for i := 0; i < leaveQueue; i++ {
		var value int
		fmt.Scan(&value)

		for key, v := range queue {
			if value == v {
				delete(queue, key)
			}
		}
	}

	sortKeys := make([]int, 0)
	for k := range queue {
		sortKeys = append(sortKeys, k)
	}

	sort.Ints(sortKeys)

	for _, v := range sortKeys {
		val, ok := queue[v]
		if ok {
			fmt.Printf("%v ", val)
			//fmt.Println(sortKeys, queue)
		} else {
			fmt.Println("não achou")
		}
	}

	fmt.Println()

}
