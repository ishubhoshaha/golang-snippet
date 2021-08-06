package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	Elements []int
}

// Add integer element to the end of the Q
func (q *Queue) Enqueue(elem int) {
	q.Elements = append(q.Elements, elem)
}

// return first element
func (q *Queue) Dequeue() (int, error){


	if len(q.Elements) == 0 {
		return 0, errors.New("Empty Q")
	}
	var firstElement int
	firstElement, q.Elements = q.Elements[0], q.Elements[1:]
	return firstElement, nil
}

// get length of Q
func (q *Queue) Size() int {
	return len(q.Elements)
}

// check q is empty or not
func (q *Queue) isEmpty() bool {
	return q.Size() == 0
}

// get top elements of Q
func (q *Queue) Peek() (int, error) {
	if q.isEmpty() {
		return -1,  errors.New("Empty Q")
	}
	return q.Elements[0], nil
}
func main()  {
	q := Queue{}
	fmt.Println(q)
	q.Enqueue(5)
	q.Enqueue(10)
	q.Dequeue()
	fmt.Println(q)
	fmt.Println(q.isEmpty())
	fmt.Println(q.Peek())
	fmt.Println(q.Size())

}

