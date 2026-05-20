package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(capacity int) *Set {
	return &Set{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (set *Set) Reserve(newCapacity int) {
	vetData := make([]int, newCapacity)

	for idx := range set.size {
		vetData[idx] = set.data[idx]
	}

	set.data = vetData
	set.capacity = newCapacity
}

func (set *Set) binarySearch(value int) int {
	low := 0
	high := set.size - 1

	for low <= high {
		mid := (low + high) / 2
		shooter := set.data[mid]
		//fmt.Println(mid, shooter)

		if value == shooter {
			return shooter
		}

		if value < shooter {
			high = mid - 1
		}

		if value > shooter {
			low = mid + 1
		}
	}

	return -1
}

func (set *Set) Contains(value int) bool {
	searchValue := set.binarySearch(value)
	if searchValue == -1 {
		return false
	}

	return true
}

func (set *Set) findInsertionIndex(value int) int {
	l := 0
	r := set.size - 1
	for l <= r {
		mid := (l + r) / 2
		if set.data[mid] > value {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}

func (set *Set) insert(value, index int) {
	if set.size == set.capacity {
		newCap := set.capacity * 2
		if newCap == 0 {
			newCap = 1
		}
		set.Reserve(newCap)
	}

	for i := set.size; i > index; i-- {
		set.data[i] = set.data[i-1]
	}
	set.data[index] = value
	set.size++
}

func (set *Set) Insert(value int) {
	if set.binarySearch(value) != -1 {
		return
	}

	set.insert(value, set.findInsertionIndex(value))
}

func (set *Set) erase(index int) {
	for i := index; i < set.size-1; i++ {
		set.data[i] = set.data[i+1]
	}

	set.size--
}

func (set *Set) Erase(value int) {
	existe := false
	for i := 0; i < set.size-1; i++ {
		if value == set.data[i] {
			existe = true
			set.erase(i)
		}
	}

	if !existe {
		fmt.Println("value not found")
	}
}

func (set *Set) String() string {
	return "[" + Join(set.data[:set.size], ", ") + "]"
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	var v *Set = NewSet(0)
	for scanner.Scan() {

		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}

		case "show":
			fmt.Println(v.String())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			v.Erase(value)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			fmt.Println(v.Contains(value))
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
