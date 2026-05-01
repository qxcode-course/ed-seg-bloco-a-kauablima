package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type MultiSet struct {
	data     []int
	size     int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet {
	return &MultiSet{
		data:     make([]int, 0, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (ms *MultiSet) Erase(value int) []int{
	newVetor := make([]int, ms.size, ms.capacity)
	copy(newVetor, ms.data)

	count := 0 
	for idx, v := range ms.data {
		if  value == v {
			count++ 
		}

		if count > 0 {
			newVetor = append(newVetor[:idx], newVetor[idx+1:]...)
			count = 0
			fmt.Println(newVetor)
			return newVetor
		}
	}

	return newVetor
}

func (ms *MultiSet) Contains(value int) bool {
	for i := 0; i < ms.size; i++ {
		v := ms.data[i]
		if v == value {
			fmt.Println("true")
			return true
		}
	}
	fmt.Println("false")
	return false
}

func (ms *MultiSet) Insert(value int) {
	ms.data = append(ms.data, value)
	ms.size = len(ms.data)
	sort.Ints(ms.data)
}

func (ms *MultiSet) String() string {
	return "[" + Join(ms.data[:ms.size], ", ") + "]"
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	// ms := NewMultiSet(0)

	var ms *MultiSet = NewMultiSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms.String())
		case "erase":
			value, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println(err)
			}
			ms.Erase(value)
		case "contains":
			value, _ := strconv.Atoi(args[1])
			ms.Contains(value)
		case "count":
			// value, _ := strconv.Atoi(args[1])
		case "unique":
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
