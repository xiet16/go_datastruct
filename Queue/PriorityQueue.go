package Queue

/*
优先队列:
没有什么是包一层解决不了的，如果有，那就再包一层
*/


type PriorityItem struct {
	value interface{}  // 数据
	Priority int //优先级
}

//队列中间的元素
func NewPriorityItem(value interface{},priority int)*PriorityItem {
	return &PriorityItem{value: value,Priority: priority}
}

//实现对比接口,对比优先级
func (x PriorityItem)Less(than Item) bool {
	return x.Priority < than.(PriorityItem).Priority
}

//基于堆的优先队列
type PriorityQueue struct {
	data *Heap
}


func NewMaxPriorityQueue() *PriorityQueue {
   return &PriorityQueue{data: NewMax()}
}

func NewMinPriorityQueue() *PriorityQueue  {
   return &PriorityQueue{data: NewMin()}
}

func (pd *PriorityQueue)Len() int {
	return pd.data.Len()
}

func (pd *PriorityQueue)Insert(el Item)  {
  pd.data.Insert(el)
}

func (pd *PriorityQueue)Extract() Item {
	return pd.data.Extract()
}


//修改优先级
func (pd *PriorityQueue)ChangePriority(value interface{},priority int)  {
	storage := NewQueue() //备份数据
	 popded := pd.Extract().(PriorityItem) //取出数据

	for value != popded.value {
		if pd.Len() ==0 {
			return
		}
		storage.EnQueue(popded)
        popded = pd.Extract().(PriorityItem)
	}

	popded.Priority = priority //修改优先级
	pd.data.Insert(popded)
    for storage.Size()>0{
    	pd.data.Insert(storage.DeQueue().(Item))
	}
}