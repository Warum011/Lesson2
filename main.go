package main

import (
	"fmt"
	"myProject/queue"
	"myProject/stack"
)

func main() {
	st := stack.NewStackOnSlice()
	st.Push(5)
	st.Push(15)
	fmt.Println(st.Print())

	q := queue.NewQueueOnSlice()
	q.PushQueue(51)
	q.PushQueue(14)
	q.PopQueue()
	fmt.Println(q.PrintQueue())

}
