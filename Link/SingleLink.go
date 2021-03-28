package Link

import (
	"fmt"
	"strings"
)

//增删改查接口
type  SingleLink interface {
    GetFirstNode() *SingleLinkNode //抓取头部节点
    InsertNodeFront(node *SingleLinkNode)  //头部插入
    InsertNodeBack(node *SingleLinkNode)  //尾部插入
	InsertNodeValueFront(dest interface{},node *SingleLinkNode)bool  //节点后插入
	InsertNodeValueBack(dest interface{},node *SingleLinkNode)bool  //节点后插入
    GetNodeAtIndex(index int) *SingleLinkNode  //根据索引获取节点
    DeleteNode(dest *SingleLinkNode) bool //删除一个节点
    DeleteAtIndex (index int) bool //删除指定位置的节点
    String() string //返回链表字符串
    FindString(query string)string //查找字符串
    Size() int //获取长度
}

//原则上，第一个节点不存储数据
type SingleLinkList struct {
	head *SingleLinkNode //链表的头指针
	length int //链表长度
}

func NewSingleLinkList() *SingleLinkList {
	head :=NewSingleLinkNode(nil)
	return &SingleLinkList{head: head,length: 0}
}


func (list *SingleLinkList)GetFirstNode() *SingleLinkNode {
	return list.head.pNext
}

func (list *SingleLinkList)InsertNodeFront(node *SingleLinkNode) {
	if list.head==nil {
		list.head.pNext = node
		node.pNext=nil
		list.length++
	}else {
		//原则上，第一个节点不存储数据
		back := list.head
		node.pNext = back.pNext
		list.head.pNext = node
		list.length ++
	}
}

func (list *SingleLinkList)InsertNodeBack(node *SingleLinkNode){
	if list.head==nil {
		list.head.pNext = node
		node.pNext=nil
		list.length++
	}else {
		back:=list.head
		for back.pNext!=nil {
			back = back.pNext
		}
		back.pNext = node
		list.length++
	}
}

func (list *SingleLinkList)GetNodeAtIndex(index int) *SingleLinkNode {
	if index >list.length-1 || index<0{
		return nil
	}

	phead := list.head
	for index >=0 {
		phead = phead.pNext
		index--
	}
	return phead
}

func (list *SingleLinkList)InsertNodeValueFront(dest interface{},node *SingleLinkNode) bool{
	phead := list.head
	isfind := false
	for phead.pNext!=nil {
		if phead.pNext.value ==dest {
			isfind =true
			break
		}
		phead=phead.pNext
	}

	if isfind {
		node.pNext = phead.pNext
		phead.pNext = node
		list.length++
		return true
	}else{
		return false
	}
}

func (list *SingleLinkList)InsertNodeValueBack(dest interface{},node *SingleLinkNode) bool{
	phead := list.head
	isfind :=false
	for phead !=nil {
		if phead.value ==dest {
           isfind = true
           break
		}
		phead = phead.pNext
	}

	if isfind {
		node.pNext = phead.pNext
		phead.pNext = node
		list.length++
		return true
	}else {
		return false
	}
}

func (list *SingleLinkList)DeleteNode(dest *SingleLinkNode) bool{
	if dest==nil {
		return false
	}

	phead := list.head
	for phead.pNext!=nil &&phead.pNext!=dest {
		phead =phead.pNext //循环
	}

	if phead.pNext == dest{
		phead.pNext = phead.pNext.pNext
		return  true
	}else {
		return false
	}
}

func (list *SingleLinkList)DeleteAtIndex (index int) bool{
	if index<0 {
		return false
	}

	phead := list.head
	for phead.pNext!=nil &&index >0 {
		phead =phead.pNext //循环
		index--
	}

	if phead.pNext !=nil{
		phead.pNext = phead.pNext.pNext
		return  true
	}else {
		return false
	}
}

func (list *SingleLinkList)String() string {
    var liststr string
    p:=list.head

	for p.pNext!=nil {
		liststr += fmt.Sprintf("%v-->",p.pNext.value)
		p=p.pNext
	}

	liststr+=fmt.Sprintf("nil")
	return liststr
}

func (list *SingleLinkList) FindString(query string)string {
	if query=="" {
        return ""
	}
    phead:=list.head.pNext
	for phead!=nil {
		if strings.Contains(phead.value.(string),query) {
           return phead.value.(string)
           break
		}
		phead =phead.pNext
	}
	return ""
}

func (list *SingleLinkList)Size() int{
	return list.length
}