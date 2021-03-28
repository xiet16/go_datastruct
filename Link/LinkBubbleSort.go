package Link

/*
单链表冒泡排序:
冒泡排序这里就不说了，可以看数组篇的
单链表的冒泡和数组的冒泡不一样的地方是不能随意交换
单链表是单向的，所以要涉及4个节点: 当前节点, 当前节点的前一个节点， 下一个节点，结束节点，
*/

func LinkBubbleSort(list *SingleLinkList)  {
	if list.head==nil || list.head.pNext==nil {
		return
	}

	phead := list.head
	cur := phead.pNext //当前循环节点,冒泡排序从第一个数开始，每次轮询都得出一个最值
	end := &SingleLinkNode{} //已经排过序的
	pre :=phead  //前一个节点

	for cur!=end{
		for cur.pNext!=nil && cur.pNext!=end {
			if cur.value.(int) > cur.pNext.value.(int) {
				//如果大于，指针往下移
				pre =cur
				cur = cur.pNext
			}else {
				//交换值
				back := cur.pNext
				cur.pNext=cur.pNext.pNext
				back.pNext=cur
				//赋值前一个节点
				pre.pNext = back
				pre=back
			}
		}
		//重新初始化
        end = cur //标记已排序过的节点
		pre = phead
        cur=phead.pNext
	}
	//return dummy
}