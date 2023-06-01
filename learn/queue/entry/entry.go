package main

import (
	"fmt"
	"test_go_mod/learn/queue"
)

func main() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Print(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Print(q.IsEmpty())
}
