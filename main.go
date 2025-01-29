package main

import (
	"fmt"
	"myProject/stack"
)

func main() {
	st := stack.NewStackOnSlice()
	st.Push(5)
	st.Push(15)
	fmt.Println(st.Print())

}
