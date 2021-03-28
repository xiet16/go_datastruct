package Queue

import "fmt"

func CircleQueueTest(){
    var queue CircleQueue
    InitQueue(&queue)
    EnQueue(&queue,1)
	EnQueue(&queue,2)
	EnQueue(&queue,3)
	EnQueue(&queue,4)

	if QueueLength(&queue)>0{
		data,_:=DeQueue(&queue)
		fmt.Println("dequeue value:",data)
	}
}

func HeapStructTest()  {
	heap:=NewMin()
	heap.Insert(MyInt(8))
	fmt.Println(heap.Extract())
	heap.Insert(MyInt(9))
	heap.Insert(MyInt(7))
	heap.Insert(MyInt(3))
	heap.Insert(MyInt(10))
	heap.Insert(MyInt(15))
	heap.Insert(MyInt(6))
	fmt.Println(heap.Extract())
}

func PriorityQueueTest(){
	h :=NewMinPriorityQueue()
	h.Insert(*NewPriorityItem(103,12))
	h.Insert(*NewPriorityItem(115,18))
	h.Insert(*NewPriorityItem(135,1))
	h.Insert(*NewPriorityItem(168,10))
	h.Insert(*NewPriorityItem(188,56))
	fmt.Println(h.Extract())
	fmt.Println(h.Extract())

}