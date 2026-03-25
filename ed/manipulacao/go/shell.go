package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	vetor := vet
	sort.Ints(vetor)

	newVetor := make([]int, 0)
	for _, v := range vetor {
		if v > 0 {
			newVetor = append(newVetor, v)
		}
	}
	return newVetor
}

func getCalmWomen(vet []int) []int {
	vetor := vet

	vetorCalm := make([]int, 0)

	for _, v := range vetor {
		if v > -10 && v < 0 {
			vetorCalm = append(vetorCalm, v)
		}
	}
	return vetorCalm
}

func sortVet(vet []int) []int {
	vetor := vet
	sort.Ints(vetor)
	return vetor
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sortStress(vet []int) []int {
	result := vet
	sort.Slice(result, func(i, j int) bool {
		return abs(result[i]) < abs(result[j])
	})

	return result
}

func reverse(vet []int) []int {
	vetorCopy := make([]int, len(vet))

	copy(vetorCopy, vet)

	slices.Reverse(vetorCopy)

	return vetorCopy
}

func unique(vet []int) []int {
	mapUnique := make(map[int]bool)
	noRepeat := make([]int, 0)

	for _, val := range vet {
		if !mapUnique[val]{
			noRepeat = append(noRepeat, val)
			mapUnique[val] = true
		}
	}

	return noRepeat
}

func repeated(vet []int) []int {
	mapRepeat := make(map[int]bool)
	repeat := make([]int, 0)

	for _, val := range vet {
		if mapRepeat[val] {
			repeat = append(repeat, val)
		} else {
			mapRepeat[val] = true
		}
	}

	return repeat
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}
