package ArrayList

import (
	"errors"
	"fmt"
)

//接口
type List interface {
   Size() int  //数组大小
   Get(index int)(interface{},error) //获取第几个元素
   Set(index int ,newval interface{}) error //修改第几个元素
   Insert(index int,newval interface{}) error //插入数据
   Append(newval interface{}) //追加
   Clear() //清空
   Delete(index int)error //删除第几个
   String() string //返回字符串
   GetIterator() Iterator //获取迭代器
}

//数据结构 这里字段是私有的
type ArrayList struct {
 dataStroe []interface{} //数组存储
 theSize int //数组大小
}

func NewArrayList() *ArrayList {
   list:=new(ArrayList)
   list.dataStroe=make([]interface{},0,10)
   list.theSize =0
   return list
}

func (list *ArrayList)Size() int  {
      return list.theSize
}

func (list *ArrayList)Get(index int)(interface{},error){
	if index<0 || index>list.theSize {
		return nil,errors.New("索引越界")
	}
	return list.dataStroe[index],nil
}

func (list *ArrayList)Append(newval interface{}){
	list.dataStroe = append(list.dataStroe,newval) //叠加数据
	list.theSize++
}

func (list *ArrayList) String() string {
   return fmt.Sprint(list.dataStroe)
}

func (list *ArrayList)Set(index int ,newval interface{}) error {
	if index<0 || index>list.theSize {
		return errors.New("索引越界")
	}
	list.dataStroe[index]=newval
	return nil
}
func (list *ArrayList)Insert(index int,newval interface{}) error {
	if index<0 || index>list.theSize {
		return errors.New("索引越界")
	}
	list.checkIsFull()//检测内存
	list.dataStroe = list.dataStroe[0:list.theSize+1] //插入数据，内存移动一位
	//从后往前移动
	for i:=list.theSize;i>index;i-- {
		list.dataStroe[i] =list.dataStroe[i-1]
	}
	list.dataStroe[index]=newval
	list.theSize++
  return  nil
}

func (list *ArrayList)checkIsFull(){
	//判断内存使用
	if list.theSize ==cap(list.dataStroe){
		//make ,中间如果是0，代表没有开辟内存。是不可以赋值的
		newDataStore:=make([]interface{},2*list.theSize,2*list.theSize) //一般是双倍内存
		copy(newDataStore,list.dataStroe)
		list.dataStroe = newDataStore
	}
}

func (list *ArrayList)Clear() {
	list.dataStroe=make([]interface{},0,10)
	list.theSize=0
}

//重新叠加
func (list *ArrayList)Delete(index int)error {
    list.dataStroe =append(list.dataStroe[:index],list.dataStroe[index+1:]...)
	list.theSize--
	return nil
}