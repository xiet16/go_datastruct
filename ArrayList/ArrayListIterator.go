package ArrayList

import "errors"

//迭代器
type Iterator interface {
    HasNext() bool  //是否有下一个

    Next() (interface{},error) //下一个

    GetIndex()int //得到索引

    Remove() //删除
}

//为什么要再定义一个
type Iterable interface {
    GetIterator() Iterator
}
type ArrayListIterator struct {
    list *ArrayList //数组指针
    currentindex int //当前索引
}

func (list *ArrayList) GetIterator() Iterator {
  it:=new(ArrayListIterator)
  it.currentindex =0
  it.list = list
	return it
}

func (it *ArrayListIterator)HasNext() bool {
   return it.currentindex<it.list.theSize
}

func (it *ArrayListIterator)Next() (interface{},error) {
	if !it.HasNext() {
		return nil,errors.New("none next value")
	}
	value,err:=it.list.Get(it.currentindex)
	if err!=nil {
		return nil, err
	}
	it.currentindex++
	return value,nil
}

func (it *ArrayListIterator)GetIndex()int{
	return it.currentindex
}

func (it *ArrayListIterator)Remove(){
	it.currentindex--
	it.list.Delete(it.currentindex)
}