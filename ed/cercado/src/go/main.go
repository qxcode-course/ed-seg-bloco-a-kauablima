package main

import (
	"bufio"
	"fmt"
	"os"
)

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	rows := len(board)
	cols := len(board[0])

	//provalmente vai ter uma variavel de identificação pra esse '0'

	var dfs func(row, col int)

	dfs = func(row, col int) {
		if row < 0 || col < 0 || rows <= row || cols <= col || board[row][col] != 'O' {
			return
		}

		board[row][col] = 'T'

		dfs(row+1, col)
		dfs(row-1, col)
		dfs(row, col+1)
		dfs(row, col-1)

	}

	for i := range board {
		for j := range board[i] {
			if board[i][0] == 'O' {
				dfs(i, 0)
			}

			if board[i][cols-1] == 'O' {
				dfs(i, cols-1)
			}

			if board[0][j] == 'O' {
				dfs(0, j)
			}
			if board[rows-1][j] == 'O' {
				dfs(rows-1, j)
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}

			if board[i][j] == 'T' {
				board[i][j] = 'O'
			}
		}
	}

}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
