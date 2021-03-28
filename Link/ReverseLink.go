package Link

/*
链表反转
*/

func (list *SingleLinkList)ReverseLink()  {
	if list.head ==nil ||list.head.pNext ==nil {
		return
	}
	var pre *SingleLinkNode
	var cur *SingleLinkNode =list.head.pNext

	//这个循环的最主要目的是给当前节点赋值上一个节点，然后继续推进
	for cur!=nil {
		curNext:= cur.pNext  //1 备份当前节点
		cur.pNext = pre      //2 当前节点等于上一节点，注意，这个pre 是上一次循环赋值的，是上一循环的当前节点
		pre =cur             //3 给pre 赋值当前节点，供下一次循环使用，第一次进来pre 是nil ,最后一个值不进来所以需要在外面赋值
		cur=curNext          //4 持续推进
	}
	list.head.pNext=pre  //最后一个值
}