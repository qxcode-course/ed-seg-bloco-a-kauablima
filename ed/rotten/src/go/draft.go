package main

import "fmt"

func search(grid [][]int) int {

    for i := range grid {
        for j := range grid[i] {
            if {
                count = -1
            }
        }
    }

    return count
}

func bfs(grid [][]int, i, j int)  {
    if i >= len(grid) ||{
        return
    }



}

func main() {
	var row, col int
    fmt.Scan(&row, &col)

    grid := make([][]int, row)

    for i := 0; i < row;i++{
        grid[i] = make([]int, col)
        for j := 0 ; j < col; j ++ {
            fmt.Scan(&grid[i][j]) 
        }
    }


    fmt.Println(search(grid))
}
