package Link

//单链表选择排序,选择排序是不需要换位置的，也可以换位置，这里就换位置吧
func LinkSelectionSort(list SingleLinkList)  {
	length := list.Size()
	if length<1 {
		return
	}

/*	phead := list.head
	pre := phead
	cur := phead.pNext
	//56789  65789 75689 85679 98567
	for cur!=nil  {
		curNext := cur.pNext
		for curNext !=nil{
			if cur.value.(int) <curNext.value.(int) {
				back:= curNext
				cur.pNext = curNext.pNext

			}
		}
	}*/
}