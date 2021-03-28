package ArrayList

type StackArrayIterator  interface {
	Clear() //清空
	Size() int //大小
	Pop() interface{} //弹出
	Push(data interface{}) //压入
	IsFull()bool //是否满
	IsEmpty()bool //是否为空

}

//引入数组和迭代器
type  StackIterator struct {
	stacklist *ArrayList
	It Iterator
}

func NewStackIterator() *StackIterator {
	stackit:= new(StackIterator)
	stackit.stacklist =NewArrayList()
	stackit.It = stackit.stacklist.GetIterator()
	return stackit
}


func (stack *StackIterator)Clear() {
	//go 自动回收内存
	stack.stacklist.Clear()
	stack.stacklist.theSize =10
}
func (stack *StackIterator)Size() int {
	return stack.stacklist.Size()
}
func (stack *StackIterator)Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	last:=stack.stacklist.dataStroe[stack.stacklist.theSize-1] //最后一个数据
	stack.stacklist.Delete(stack.stacklist.theSize-1)
	return last
}
func (stack *StackIterator)Push(data interface{}) {
	if stack.IsFull() {
		return
	}

	stack.stacklist.Append(data)
}
func (stack *StackIterator)IsFull()bool {
	return  stack.stacklist.theSize >=10
}
func (stack *StackIterator)IsEmpty()bool {
	return stack.stacklist.theSize ==0
}