package godata

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue(3)
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	fmt.Println(queue.Top())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	// nil
	fmt.Println(queue.Pop())
	fmt.Println(queue.Length())
	fmt.Println(queue.val)
	fmt.Println("top",queue.Top())

}
