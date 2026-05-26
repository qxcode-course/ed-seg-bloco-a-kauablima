package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {
	var str []string
	for _, v := range vet {
		str = append(str, strconv.Itoa(v))
	}

	strPrint := "["
	for idx, v := range str {
		strPrint += v
		if idx != len(str)-1 {
			strPrint += ", "
		}
	}
	strPrint += "]"
	return strPrint
}

func tostrrev(vet []int) string {
	n := len(vet)
	if n == 0 {
		return "[]"
	}

	strVet := make([]string, 0, n)

	for i := n - 1; i >= 0; i-- {
		strVet = append(strVet, strconv.Itoa(vet[i]))
	}

	return "[" + strings.Join(strVet, ", ") + "]"
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	if len(vet) <= 1 {
		return
	}

	idx := len(vet) - 1

	vet[0], vet[idx] = vet[idx], vet[0]
	reverse(vet[1:idx])
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	if len(vet) == 0 {
		return 0
	}
	return vet[0] + sum(vet[1:])
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	if len(vet) == 0 {
		return 1 // Elemento neutro da multiplicação
	}
	return vet[0] * mult(vet[1:])
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	if len(vet) == 0 {
		return -1
	}

	var rec func(v []int, currentIdx int) (int, int)
	rec = func(v []int, currentIdx int) (int, int) {
		if len(v) == 1 {
			return v[0], currentIdx
		}

		minVal, minIdx := rec(v[1:], currentIdx+1)

		if v[0] <= minVal {
			return v[0], currentIdx
		}
		return minVal, minIdx
	}

	_, index := rec(vet, 0)
	return index
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
