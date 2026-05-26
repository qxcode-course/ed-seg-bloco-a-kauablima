package main

import (
	"bufio"

	"fmt"
	"os"
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
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (ms *MultiSet) expand() {
	newCapacity := ms.size * 2
	vetData := make([]int, newCapacity)

	for i := range ms.size {
		vetData[i] = ms.data[i]
	}

	ms.data = vetData
	ms.capacity = newCapacity
}

// func (ms *MultiSet) search(value int) (bool, int) {
// 	low := 0
// 	high := ms.size - 1

// 	for low <= high {
// 		mid := (low + high) / 2

// 		if ms.data[mid] == value {
// 			return true, mid
// 		}

// 		if ms.data[mid] < value {
// 			low = mid + 1
// 		}

// 		if ms.data[mid] > value {
// 			high = mid - 1
// 		}
// 	}

// 	return false, low
// }

// func (ms *MultiSet) findValue(value int) int {
// 	low := 0
// 	high := ms.size - 1

// 	for low <= high {
// 		mid := (low + high) / 2

// 		if ms.data[mid] == value {
// 			return mid
// 		}

// 		if ms.data[mid] < value {
// 			low = mid + 1
// 		}

// 		if ms.data[mid] > value {
// 			high = mid - 1
// 		}
// 	}

// 	return -1
// }

func (ms *MultiSet) QueryValueIdx(value int) int {
	low := 0
	high := ms.size - 1

	for low <= high {
		mid := (low + high) / 2

		if ms.data[mid] == value {
			return mid
		}

		if ms.data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func (ms *MultiSet) erase(index int) {
	for i := index; i < ms.size; i++ {
		ms.data[i] = ms.data[i+1]
	}

	ms.size--
}

func (ms *MultiSet) Erase(value int) error {
	if ms.QueryValueIdx(value) == -1 {
		return fmt.Errorf("value not found")
	}

	ms.erase(ms.QueryValueIdx(value))
	return nil
}

func (ms *MultiSet) QueryIndex(value int) int {
	low := 0
	high := ms.size - 1

	for low <= high {
		mid := (low + high) / 2

		if ms.data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}

func (ms *MultiSet) Insert(value int) {
	if ms.size == ms.capacity {
		ms.expand()
	}

	ms.insert(value, ms.QueryIndex(value))
}

func (ms *MultiSet) insert(value, index int) {
	for i := ms.size; i > index; i-- {
		ms.data[i] = ms.data[i-1]
	}
	ms.data[index] = value
	ms.size++
}

func (ms *MultiSet) Count(value int) int {
	count := 0
	for i := 0; i < ms.size; i++ {
		if ms.data[i] == value {
			count++
		}
	}

	if count != 0 {
		return count
	}

	return 0
}

func (ms *MultiSet) Unique() int {
	if ms.size == 0 {
		return 0
	}

	count := 1
	for i := range ms.size - 1 {
		if ms.data[i] != ms.data[i+1] {
			count++
		}
	}

	return count
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

func (ms *MultiSet) Clear() {
	for idx, val := range ms.data {
		if val == 0 {
			ms.data[idx] = 0
		}
	}
	ms.size = 0
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
			value, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println(err)
			}
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms.String())
		case "erase":
			value, _ := strconv.Atoi(args[1])
			err := ms.Erase(value)
			if err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			ms.Contains(value)
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
