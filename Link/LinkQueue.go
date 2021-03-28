package Link

type LinkQueue struct {
	front *Node //头部节点
	rear *Node //尾部节点
}

type ILinkQueue interface {
    Length()int
	Enqueue(value interface{})
	Dequeue()interface{}
}

func NewLinkQueue() *LinkQueue{
	return &LinkQueue{}
}

func (q *LinkQueue) Length()int {
	length:=0
	pnext :=q.front
	for pnext.pNext!=nil {
		pnext =pnext.pNext
		length ++
	}
	return length
}

//头部插入
func (q *LinkQueue)Enqueue(value interface{})  {
	newNode := &Node{data: value,pNext: nil}
	if q.front==nil {
		q.front = newNode
		q.rear = newNode
	}else {
		q.rear.pNext =newNode
		q.rear=q.rear.pNext
	}
}

func (q *LinkQueue)Dequeue()interface{}  {
	if q.front ==nil {
		return  nil
	}
	res := q.front.data
	//只有一个元素
	if q.front == q.rear {
		q.front =nil
		q.rear =nil
	}else {
		q.front = q.front.pNext
	}
	return res
}