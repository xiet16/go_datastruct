package Link

import "fmt"

func LinkStackTest()  {
	/*
	stack := NewStack()
	for i :=0;i<10;i++ {
		stack.Push(i)
	}

	for data:= stack.Pop();data!=nil;data=stack.Pop() {
		fmt.Println(data)
	}
	*/
	//DoubleLinkTest()
	SingleLinkTest()

}

func SingleLinkTest()  {
	linkList :=NewSingleLinkList()
	linkList.InsertNodeFront(&SingleLinkNode{value: 1})
	fmt.Println(linkList.String())
	linkList.InsertNodeBack(&SingleLinkNode{value: 2})
	fmt.Println(linkList.String())
	node3:=&SingleLinkNode{value: 3}
	linkList.InsertNodeBack(node3)
	fmt.Println(linkList.String())
	fmt.Println(linkList.String())
	node4:=&SingleLinkNode{value: 4}
	linkList.InsertNodeFront(node4)
	fmt.Println(linkList.String())
	linkList.DeleteNode(node3)
	fmt.Println(linkList.String())
	linkList.InsertNodeValueBack(2,&SingleLinkNode{value: 5})
	fmt.Println(linkList.String())
	linkList.InsertNodeValueFront(1,&SingleLinkNode{value: 6})
	fmt.Println(linkList.String())
	linkList.InsertNodeValueFront(4,&SingleLinkNode{value: 0})
	fmt.Println(linkList.String())
	fmt.Println(linkList.GetNodeAtIndex(2))
	//linkList.ReverseLink()
	//fmt.Println(linkList.String())
	LinkBubbleSort(linkList)
	fmt.Println(linkList.String())
}

func DoubleLinkTest()  {
	linkList := NewDoubleLinkList()
	linkList.InsertHead(&DoubleLinkNode{value: 1})
	fmt.Println(linkList.String())
	linkList.InsertBack(&DoubleLinkNode{value: 2})
	fmt.Println(linkList.String())
	node3:=&DoubleLinkNode{value: 3}
	linkList.InsertBack(node3)
	fmt.Println(linkList.String())
	linkList.InsertValueBack(node3,&DoubleLinkNode{value: 8})
	fmt.Println(linkList.String())
	node4:=&DoubleLinkNode{value: 4}
	linkList.InsertHead(node4)
	fmt.Println(linkList.String())
	linkList.InsertValueHead(node4,&DoubleLinkNode{value: 9})
	fmt.Println(linkList.String())
	linkList.InsertBack(&DoubleLinkNode{value: 5})
	fmt.Println(linkList.String())
	linkList.InsertBack(&DoubleLinkNode{value: 6})
	fmt.Println(linkList.String())
	linkList.InsertHead(&DoubleLinkNode{value: 0})
	fmt.Println(linkList.String())


}