package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type Node struct {
	Value int
	next *Node
	prev *Node
	root *Node
}

func (n *Node) Next() *Node {
	if  n.next == n.root { return nil } 

	return n.next
}

func (n *Node) Prev() *Node {
	if n.prev == n.root { return nil } 

	return n.prev
}

type LList struct {
	root *Node
	size int 
}

func NewLList() *LList {
	list := &LList{
		root: &Node{},
		size: 0,
	}

	list.root.next = list.root
	list.root.prev = list.root
	list.root.root = list.root

	return list
}

func (l *LList) Size() int {
	return l.size

}

func (l *LList) Clear() {
	l.root.next = l.root
	l.root.prev = l.root
	l.size = 0
}

func (l *LList) PushFront(val int)  {
	newNode := &Node{Value: val, next: l.root.next, prev: l.root, root: l.root}

	l.root.next.prev = newNode
	l.root.next = newNode

	l.size++
}
func (l *LList) PushBack(val int) {
	newNode := &Node{Value: val, next: l.root, prev: l.root.prev, root: l.root}

	l.root.prev.next = newNode
	l.root.prev = newNode

	l.size++
}
func (l *LList) PopFront() {
	l.root.next.prev = nil
	l.root.next = l.root.next.next
	l.root.next.next = nil
}

func (l *LList) PopBack() {
	l.root.prev.next = nil
	l.root.prev = l.root.prev.prev
	l.root.prev.prev = nil
}

func (l *LList) Front() *Node {
	if l.root.Next() == nil {
		return nil
	}

	return l.root.next
}

func (l *LList) Back() *Node  {
	if l.root.Prev() == nil {
		return nil
	}

	return l.root.prev
}
func (l *LList) Search(value int) *Node{
	for node := l.root.Next(); node != nil;node = node.Next(){
		if value == node.Value {
			return node
		}
	}

	return nil 
}

func (l *LList) Insert(node *Node, value int) {
	if node == nil {
		return 
	}	

	newNode := &Node{Value: value, next: node, prev: node.prev, root: node.root}

	previous := node.prev
	node.prev = newNode
	previous.next = newNode
	l.size++
}

func (l *LList) Remove(node *Node) *Node { 
	NNode := node.next
	PNode := node.prev

	PNode.next = NNode
	node.prev = nil
	node.next =nil
	l.size--
	return NNode
}

func (l *LList) String() string {
	str := "["

	for node := l.root.next; node != l.root; node = node.next {
		value := fmt.Sprint(node.Value)
		str += value

		if node.next != l.root {
			str += ", "
		}
	}

	str += "]"
	return str
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
 	ll := NewLList()

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
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.Value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
