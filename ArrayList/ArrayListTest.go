package ArrayList

import "fmt"

func Test1()  {
	fmt.Println("welcome to datastruct")
	list:=NewArrayList()
	list.Append(1)
	list.Append(5)
	list.Append(8)
	list.Insert(2,66)

	for i:=0;i<10;i++ {
		list.Insert(i,77)
	}
	fmt.Println(list.String())
}

func IteratorTest() {
	fmt.Println("welcome to datastruct")
	list:=NewArrayList()
	list.Append(1)
	list.Append(5)
	list.Append(8)
	list.Insert(2,66)

	for it:=list.GetIterator();it.HasNext();{
		item,_ :=it.Next()
		fmt.Println(item)
		if item ==5{
			it.Remove()
		}
	}
	fmt.Println(list.String())
}

//引用数组测试
func StackArrayTest()  {
	stack := NewArrayListStack()
	stack.Push("a1")
	stack.Push("a2")
	stack.Push("a3")
	stack.Push("a4")

	for i:=0;i<4;i++{
		fmt.Println("stack data is: ",stack.Pop())
	}
}

//引入迭代器测试
func StackArrayIteratorTest()  {
	stack := NewStackIterator()
	stack.Push("a1")
	stack.Push("a2")
	stack.Push("a3")
	stack.Push("a4")

	for it:=stack.It;it.HasNext();{
		fmt.Println("stack data is: ",stack.Pop())
	}
}