package StackArray

import (
	"fmt"
)

func StackArrayTest()  {
	stack := NewStack()
	stack.Push("a1")
	stack.Push("a2")
	stack.Push("a3")
	stack.Push("a4")

	for i:=0;!stack.IsEmpty();i++{
		fmt.Println("stack data",i+1," is: ",stack.Pop())
	}
}
