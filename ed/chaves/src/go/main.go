package main

import "fmt"

func main() {
	queue := NewQueue[string]()

	for i := 'A'; i <= 'P'; i++ {
		queue.Enqueue(string(i))
	}

	for i := 0; i < 15; i++ {
		t1 := queue.Dequeue()
		t2 := queue.Dequeue()

		var g1, g2 int
		fmt.Scan(&g1, &g2)

		if g1 > g2 {
			queue.Enqueue(t1)
		} else {
			queue.Enqueue(t2)
		}
	}
	
	fmt.Println(queue.Dequeue())
}
