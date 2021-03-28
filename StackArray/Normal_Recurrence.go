package StackArray

import (
	"fmt"
)

//1+2+3+n
//recurrence 递归
//栈是深度遍历
func Add(num int) int {
	if num ==0 {
		return 0
	}
	return num + Add(num-1)
}

//栈中的递归

func StackRecurrence()  {
	stack:=NewStack()
	stack.Push(5)
	last:=0

	for !stack.IsEmpty() {
		data :=stack.Pop()

		if data ==0{
			last +=0
		}else {
			last +=data.(int)
			stack.Push(data.(int)-1)
		}
	}
	fmt.Println(last)
}



//f(n) = f(n-1)+f(n-2)
//使用方法实现
func RecFuncAdd(num int)int  {
	if num ==1 ||num==2 {
		return 1
	}else {
		return RecFuncAdd(num-1) + RecFuncAdd(num-2)
	}
}

func RecTest()  {
	fmt.Println(RecStackAdd(6))
}

//使用栈实现
func RecStackAdd(num int) int {
	stack := NewStack()
    stack.Push(num)
	last:=0
	for !stack.IsEmpty() {
		data :=stack.Pop().(int)
		if data==1||data==2 {
			last +=1
		}else {
			stack.Push(data-2)
			stack.Push(data-1)
		}
	}
	return last
}
