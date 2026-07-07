package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	text := scanner.Text()
	vetRune := []rune{}

	for _, c := range text {

		switch c {
		case '(':
			vetRune = append(vetRune, ')')
		case '[':
			vetRune = append(vetRune, ']')
		case ')', ']':
			if len(vetRune) == 0 || vetRune[len(vetRune)-1] != c {
				fmt.Println("nao balanceado")
				return
			}

			vetRune = vetRune[:len(vetRune)-1]
		}
	}
	if len(vetRune) != 0 {
		fmt.Println("nao balanceado")
        return
	}

	fmt.Println("balanceado")
}
