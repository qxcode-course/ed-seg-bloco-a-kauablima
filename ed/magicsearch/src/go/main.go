package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
//buscabinaria

func MagicSearch(slice []int, value int) int {
	high := len(slice) -1 
	low := 0

	for high >= low {
		mid := (high +low)/2  

		if value >= slice[mid]{
			low = mid + 1
		}

		if value < slice[mid]{
			high = mid -1 
		}

	}

	if low > 0 && slice[low-1] == value {
		return low -1
	}

	return low
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice, value)
	fmt.Println(result)
}
