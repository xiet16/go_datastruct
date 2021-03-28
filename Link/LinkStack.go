package Link

type Node struct {
	data interface{}
	pNext *Node
}

type LinkStack interface {
	IsEmpty() bool
	Push(data interface{})
	Pop() interface{}
	Length() int
}

func NewStack() *Node {
	return &Node{} //返回一个节点的指针
}

func (n *Node)IsEmpty() bool{
	return n.pNext ==nil
}

//头部插入，头部删除
func (n *Node)Push(data interface{}){
    newNode := &Node{data: data}
    newNode.pNext=n.pNext //这一步我理解
    n.pNext = newNode //这一步是为什么？
}
func (n *Node)Pop() interface{}{
  if n.IsEmpty()==true{
  	return nil
  }
  value := n.pNext.data //后进先出，都是操作头部
  n.pNext = n.pNext.pNext //删除

  return  value
}

func (n *Node)Length() int{
  next:=n
  length:=0
	for next.pNext !=nil {
		next = next.pNext
		length ++
	}
	return length
}