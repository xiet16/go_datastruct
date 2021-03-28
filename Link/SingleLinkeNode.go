package Link

//单链表节点
type SingleLinkNode struct {
	value interface{}
	pNext *SingleLinkNode
}

func NewSingleLinkNode(data interface{})*SingleLinkNode{
	return &SingleLinkNode{value: data,pNext: nil}
}

//
func (node *SingleLinkNode)Value()interface{}  {
	return node.value
}

//返回下一节点
func (node *SingleLinkNode)PNext() *SingleLinkNode {
   return node.pNext
}
