package Link

/*
求中间值，或者说求n份之一值
*/

func (list *SingleLinkList)GetMiddle() *SingleLinkNode {
	if list.head.pNext ==nil{
		return nil
	}else {
		phead1:=list.head
		phead2:=list.head

		for phead2!=nil&&phead2.pNext!=nil {
            phead2 = phead2.pNext.pNext
            phead1 =phead1.pNext
		}
		return phead1
	}
}