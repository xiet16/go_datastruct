package Link

import "fmt"

//双链表基本结构
type DoubleLinkList struct {
	head *DoubleLinkNode
	length int
}

//新建双链表
func NewDoubleLinkList() *DoubleLinkList {
	head:=NewDoubleLinkNode(nil)
	return &DoubleLinkList{head: head,length: 0}
}

func (dlist *DoubleLinkList)Size() int {
	return dlist.length
}

func (dlist *DoubleLinkList)GetFirstNode() *DoubleLinkNode {
	return dlist.head.next
}

func (dlist *DoubleLinkList)InsertHead(node *DoubleLinkNode) {
	phead := dlist.head
	if phead.next ==nil {
		node.next =nil
		phead.next =node
		node.prev =phead
		dlist.length ++
	}else {
		pnext :=phead.next
		node.next = pnext
		pnext.prev = node
		node.prev = phead
		phead.next = node

		dlist.length ++
	}

}


func (dlist *DoubleLinkList)InsertBack(node *DoubleLinkNode) {
	phead := dlist.head
	if phead.next ==nil {
		node.next =nil
		phead.next =node
		node.prev =phead
		dlist.length ++
	}else {
		for phead.next!=nil {
            phead = phead.next
		}
		phead.next=node
		node.prev = phead
		dlist.length++
	}
}


func (dlist *DoubleLinkList)InsertValueHead(dest,node *DoubleLinkNode) bool {
	phead := dlist.head
	for phead.next !=nil && phead.next!= dest{
       phead = phead.next
	}

	if phead.next==dest {
		pnext :=phead.next
		node.next = pnext
		pnext.prev = node
		node.prev = phead
		phead.next = node

		dlist.length ++
		return true
	}else {
		return false
	}
}


func (dlist *DoubleLinkList)InsertValueBack(dest,node *DoubleLinkNode) bool {
	phead := dlist.head
	for phead != dest{
		phead = phead.next
	}

	if phead==dest {
		pnext :=phead.next
		node.next = pnext
		if pnext !=nil {
			pnext.prev = node
		}
		node.prev = phead
		phead.next = node

		dlist.length ++
		return true
	}else {
		return false
	}
}

func (dlist *DoubleLinkList)String() string {
   var listStr string

   phead := dlist.head
	for phead.next!=nil {
		listStr+=fmt.Sprintf("%v-->",phead.next.value)
		phead = phead.next
	}
	listStr += fmt.Sprintf("nil")

    return listStr
}

func (dlist *DoubleLinkList)GetNodeAtIndex(index int) *DoubleLinkNode {
	if index>dlist.length-1||index<0 {
		return nil
	}
	phead:=dlist.head
	for index >-1 {
		phead =phead.next
		index--
	}
	return phead
}

func (dlist *DoubleLinkList)DeleteNode(node *DoubleLinkNode) bool {
	if node ==nil {
		return false
	}
	phead:=dlist.head
	for phead.next!=nil&& phead.next !=node {
		phead = phead.next
	}

	if phead.next==node {
		if phead.next.next!=nil {
			phead.next.next.prev = phead
		}
		phead.next = phead.next.next
		dlist.length--
		return true
	}
	return false
}

func (dlist *DoubleLinkList)DeleteNodeAtIndex(index int) bool {
	if index>dlist.length-1||index<0 {
		return false
	}
	phead :=dlist.head
	for index >-1{
		phead =phead.next
		index--
	}

	if phead.next!=nil {
		phead.next.prev = phead.prev
	}
	phead.prev.next = phead.next
	dlist.length--
	return true
}