package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type Dlista struct {
	head *Node
}

func NewDlist() *Dlista {
	list := &Dlista{}
	list.head = &Node{}

	list.head.next = list.head
	list.head.prev = list.head

	return list
}

func (l *Dlista) Size() int {
	count := 0
	for node := l.head.next; node != l.head; node = node.next {
		count++
	}

	return count
}

func (l *Dlista) Clear() {
	l.head.prev = l.head
	l.head.next = l.head
}

func (l *Dlista) PushFront(val int) {
	newNode := &Node{value: val}

	newNode.prev = l.head
	newNode.next = l.head.next

	l.head.next.prev = newNode
	l.head.next = newNode
}

func (l *Dlista) PushBack(value int) {
	newNode := &Node{value: value}

	newNode.prev = l.head.prev
	newNode.next = l.head

	l.head.prev.next = newNode
	l.head.prev = newNode
}

func (l *Dlista) PopBack() {
	if l.Size() == 0 {
		return
	}
	nodetoPop := l.head.prev
	newPrev := nodetoPop.prev

	l.head.prev = newPrev
	newPrev.next = l.head

	nodetoPop.next = nil
	nodetoPop.prev = nil
}

func (l *Dlista) PopFront() {
	if l.Size() == 0 {
		return
	}
	nodetoPop := l.head.next
	newnext := nodetoPop.next

	l.head.next = newnext
	newnext.prev = l.head

	nodetoPop.next = nil
	nodetoPop.prev = nil
}

func (l *Dlista) String() string {
	str := "["

	for node := l.head.next; node != l.head; node = node.next {
		value := fmt.Sprint(node.value)
		str += value

		if node.next != l.head {
			str += ", "
		}
	}

	str += "]"
	return str
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewDlist()

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
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
