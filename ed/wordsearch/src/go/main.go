package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	rows := len(grid)
	cols := len(grid[0])

	var backtracking func(r, c, k int) bool

	backtracking = func(r, c, k int) bool {
		if k == len(word) {
			return true
		}

		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != word[k] {
			return false
		}

		temp := grid[r][c]
		grid[r][c] = '#'

		found := backtracking(r-1, c, k+1) || 
				 backtracking(r+1, c, k+1) ||
				 backtracking(r, c-1, k+1) || 
				 backtracking(r, c+1, k+1)
				 

		grid[r][c] = temp

		return found

	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == word[0] && backtracking(i, j, 0) {
				return true
			}
		}
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)

	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}

	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
