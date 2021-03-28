package Link

type DoubleLinkNode struct {
	value interface{}
	prev *DoubleLinkNode
	next *DoubleLinkNode
}

func NewDoubleLinkNode(value interface{}) *DoubleLinkNode {
	return &DoubleLinkNode{value:value,prev: nil,next: nil}
}

func (node *DoubleLinkNode)Value()interface{}{
  return node.value
}

//返回上一个节点
func (node *DoubleLinkNode)Prev()  *DoubleLinkNode {
  return node.prev
}

//返回下一个节点
func (node *DoubleLinkNode)PNext() *DoubleLinkNode{
	return node.next
}