package Queue

import (
	"sync"
)

/*
堆原理
*/

//接口
type Item interface {
	Less(than Item) bool //比大小
}

// 最小堆  最大堆
type Heap struct {
	lock *sync.Mutex //线程安全
	data []Item //数组
	min bool //最小堆还是最大堆
}

//标准堆
func NewHeap()*Heap {
	return &Heap{new(sync.Mutex),make([]Item,0),true}
}

//最小dui
func NewMin()*Heap {
	return &Heap{new(sync.Mutex),make([]Item,0),true}
}

//最大堆
func NewMax()*Heap {
	return &Heap{new(sync.Mutex),make([]Item,0),false}
}

//是否为空
func (h *Heap)IsEmpty() bool {
	return len(h.data)==0
}

//获取长度
func (h *Heap)Len()int  {
	return len(h.data)
}

//获取数据
func (h *Heap)Get(index int)Item {
	return h.data[index]
}

//插入数据
func (h *Heap)Insert(it Item)  {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.data =append(h.data,it)
	h.SiftUp() //重新升一下序
	return
}

//根据类型返回比大小
func (h *Heap)Less(a,b Item) bool{
	if h.min {
	   return a.Less(b)
	}else {
	   return b.Less(a)
	}
}

//压缩 相当于弹出一个
func (h *Heap)Extract()Item{
	h.lock.Lock()
	defer h.lock.Unlock()
	if h.Len() ==0 {
		return nil //长度为0 不需要
	}

	el := h.data[0]
	last:= h.data[h.Len()-1]
	if h.Len()==1 {
		h.data =nil
		return el
	}

	//这里为什么这么写呢,奇怪
	h.data =append([]Item{last},h.data[1:h.Len()-1]...)
	h.SiftDown()
	return el
}

//取出极大值
func (h *Heap)SiftUp() {
	//堆排序的循环过程  n , 2n+1
	for i,parent :=h.Len()-1,h.Len()-1;i>0;i=parent{
		parent =i/2
		if h.Less(h.Get(i),h.Get(parent)) {
			h.data[parent],h.data[i] = h.data[i],h.data[parent]
		}else {
			break
		}

	}
}


//取出极小值
func (h *Heap)SiftDown() {
	//堆排序的循环过程  n , 2n+1,2n+2
	for i,child :=0,1;i<h.Len()&&i*2+1<h.Len();i=child{
		child =i*2+1
		if child+1<h.Len()-1 && h.Less(h.Get(child+1),h.Get(child)) {
			child++
		}

		if h.Less(h.Get(i),h.Get(child)) {
			break
		}

		h.data[i],h.data[child] = h.data[child],h.data[i] //交换数据
	}
}