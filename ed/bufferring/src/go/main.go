package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Deque struct {
	data     []int
	front    int
	size     int
	capacity int
}

func (b *Deque) String() string {
	result := []string{}
	for i := range b.size {
		val := b.data[(b.front+i)%b.capacity]
		result = append(result, fmt.Sprint(val))
	}
	return "[" + strings.Join(result, ", ") + "]"
}

func (b *Deque) Debug() string {
	result := make([]string, b.capacity)
	for i, _ := range result {
		result[i] = " _"
		if i == b.front {
			result[i] = ">_"
		}
	}
	for i := range b.size {
		index := (b.front + i) % b.capacity
		val := b.data[index]
		prefix := " "
		if i == 0 {
			prefix = ">"
		}
		result[index] = fmt.Sprintf("%s%d", prefix, val)
	}
	return strings.Join(result, " |")
}

func (d *Deque) resize(newcap int) {
	newDq := make([]int, newcap)
	for interations, i := 0, d.front; interations < d.size; interations, i = interations+1, (i+1)%d.capacity {
		newDq[interations] = d.data[i]
	}
	d.data = newDq
	d.capacity = newcap
	d.front = 0

}

func (d *Deque) Len() int {
	return d.size
}

func (d *Deque) PushBack(v int) {
	if d.size == d.capacity {
		if d.size == 0 {
			d.resize(d.size + 1)
		} else {
			d.resize(d.size * 2)
		}
	}

	d.data[(d.front+d.size)%d.capacity] = v
	d.size++
}
func (d *Deque) PushFront(v int) {
	if d.size == d.capacity {
		if d.size == 0 {
			d.resize(d.size + 1)
		} else {
			d.resize(d.size * 2)
		}
	}

	d.front = (d.front - 1 + d.capacity) % d.capacity
	d.data[d.front] = v
	d.size++
}

func (d *Deque) PopBack() error {

	if d.size == 0 {
		return fmt.Errorf("fail: buffer vazio")
	}
	d.size--

	return nil
}
func (d *Deque) PopFront() error {
	if d.size == 0 {
		return fmt.Errorf("fail: buffer vazio")
	}

	d.front = (d.front + 1) % d.capacity
	d.size--

	return nil
}

func (d *Deque) Front() (int, error) {
	if d.size == 0 {
		return 0, fmt.Errorf("fail: buffer vazio")
	}
	return d.data[d.front], nil
}

func (d *Deque) Back() (int, error) {
	if d.size == 0 {
		return 0, fmt.Errorf("fail: buffer vazio")
	}
	return d.data[(d.front+d.size-1+d.capacity)%d.capacity], nil
}

func (d *Deque) Clear() {
	d.size = 0
	d.front = 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buf := &Deque{data: make([]int, 4), capacity: 4}

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(buf.String())
		case "debug":
			fmt.Println(buf.Debug())
		case "size":
			fmt.Println(buf.Len())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				buf.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				buf.PushFront(num)
			}
		case "pop_back":
			if err := buf.PopBack(); err != nil {
				fmt.Println(err)
			}
		case "pop_front":
			if err := buf.PopFront(); err != nil {
				fmt.Println(err)
			}
		case "front":
			if val, err := buf.Front(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(val)
			}
		case "back":
			if val, err := buf.Back(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(val)
			}
		case "clear":
			buf.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
