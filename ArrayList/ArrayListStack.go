package ArrayList

type StackArray interface {
	Clear() //清空
	Size() int //大小
	Pop() interface{} //弹出
	Push(data interface{}) //压入
	IsFull()bool //是否满
	IsEmpty()bool //是否为空

}

type  Stack struct {
	stackArray *ArrayList //为什么要这么写，代码复用
	capsize int //最大范围
}

func NewArrayListStack() *Stack {
	stack:= new(Stack)
	stack.stackArray =NewArrayList()
	stack.capsize =10
	return stack
}

func (stack *Stack)Clear() {
	//go 自动回收内存
	stack.stackArray.Clear()
	stack.capsize =10
}
func (stack *Stack)Size() int {
	return stack.stackArray.Size()
}
func (stack *Stack)Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	last:=stack.stackArray.dataStroe[stack.stackArray.theSize-1] //最后一个数据
	stack.stackArray.Delete(stack.stackArray.theSize-1)
	return last
}
func (stack *Stack)Push(data interface{}) {
	if stack.IsFull() {
		return
	}

	stack.stackArray.Append(data)
	//stack.stackArray.theSize++  //这里不自增吗
}
func (stack *Stack)IsFull()bool {
	return  stack.stackArray.theSize == stack.capsize
}
func (stack *Stack)IsEmpty()bool {
	return stack.stackArray.theSize ==0
}