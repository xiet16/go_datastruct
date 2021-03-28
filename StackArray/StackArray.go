package StackArray

/*
栈是一种数据结构,是只能在一端插入和删除数据的特殊线性表
先进后出
入栈和出栈都在栈顶进行
栈最主要的是保存了函数调用相关信息:其中最主要的是
1、函数的返回地址和参数
2、临时变量
*/

type StackArray interface {
   Clear() //清空
   Size() int //大小
   Pop() interface{} //弹出
   Push(data interface{}) //压入
   IsFull()bool //是否满
   IsEmpty()bool //是否为空

}

type  Stack struct {
   dataSource []interface{}
   capsize int //最大范围
   currentsize int //当前大小
}

func NewStack() *Stack {
	stack:= new(Stack)
	stack.dataSource =make([]interface{},0,10)
	stack.capsize =10
	stack.currentsize =0
	return stack
}

func (stack *Stack)Clear() {
	//go 自动回收内存
	stack.dataSource =make([]interface{},0,10)
	stack.capsize =10
	stack.currentsize =0
}


func (stack *Stack)Size() int {
  return stack.currentsize
}

//入栈
func (stack *Stack)Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	last:=stack.dataSource[stack.currentsize-1] //最后一个数据
	stack.dataSource=stack.dataSource[:stack.currentsize-1]
	stack.currentsize--
	return last
}

//出栈
func (stack *Stack)Push(data interface{}) {
	if stack.IsFull() {
		return
	}

	stack.dataSource =append(stack.dataSource,data)
	stack.currentsize++
}


func (stack *Stack)IsFull()bool {
  return  stack.currentsize == stack.capsize
}


func (stack *Stack)IsEmpty()bool {
  return stack.currentsize ==0
}