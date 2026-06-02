package main
import "fmt"


func search(vet []int, idx, curSum int, sum int, count int) bool {
    if curSum == sum {
        return true
    }

    if count == len(vet) || curSum > sum {
        return false 
    }

    nextIdx := (idx+1)%len(vet)

    return search(vet, nextIdx, curSum+vet[idx], sum, count+1 ) ||
            search(vet, nextIdx, curSum, sum, count+1)
}

func main() {
    var num, sum int
    fmt.Scan(&num, &sum)

    vet := make([]int, num) 
    for i := range num {
        fmt.Scan(&vet[i])
    }

    fmt.Println(search(vet, 0, 0, sum, 0))
}
