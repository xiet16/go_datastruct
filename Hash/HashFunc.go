package Hash

import (
	"crypto/sha256"
	"errors"
)

/*
哈希函数
哈希的原理是:任何一个字符串都能转换为一个整形，同一字符串对应的哈希值一定相同，不同字符串对应的哈希值一定不同，同一哈希值对应的字符串不一定相同

哈希表:
不适合频繁插入删除
可以实现快速查找
*/

const (
	Deleted = iota  //数据已经被删除
	MintableSize =100 //哈希表大小
	legimate =iota //已经存在的合法数据
	Empty = iota  //数据为空
)

//自定义哈希
func MySha(str interface{},tableSize int) int {
	var hashVal int= 0
	var chars []byte
	if strings,ok:=str.(string);ok {
		chars = []byte(strings)
	}

	for _,v := range chars{
		hashVal = (hashVal<<17|123&234)+int(v)
	}

	return hashVal%MintableSize
}

//sha256
func MySha256(str interface{},tableSize int) int {
   shaobj:=sha256.New()
   shaobj.Write([]byte(str.(string))) //哈希
   mybytes :=shaobj.Sum(nil)

	var hashVal int= 0
	for _,v := range mybytes{
		hashVal = (hashVal<<17|123&234)+int(v)
	}

	return hashVal%MintableSize
}

//定义哈希函数指针
type HashFunc func(data interface{},tableSize int) int //函数指针


//哈希表--> 哈希结构体集合-->哈希结构体
type HashEntry struct {
	data interface{} //数据
	kind int //类型
}

type HashTable struct {
	tableSize int //哈希表大小
	theCells [] *HashEntry  //数组，每一个元素就是指针，指向哈希结构
	hashFunc HashFunc    //调用哈希函数
}

type HashTableGo interface {
	Find(data interface{}) int
    Insert (data interface{})
	Empty()
	GetValue(index int)
}

func NewHashTable(size int,hash HashFunc) (*HashTable,error) {
	if size <MintableSize {
		return nil,errors.New("哈希表太小")
	}

	if hash == nil {
		return nil,errors.New("没有哈希函数")
	}

	//设置哈希表并且返回
	hashTable := &HashTable{ size, make([]*HashEntry,size), hash}
	for i:=0;i<hashTable.tableSize;i++ {
		hashTable.theCells[i] =new(HashEntry)
		hashTable.theCells[i].data =nil
		hashTable.theCells[i].kind = Empty
	}

	return hashTable,nil
}

func (h *HashTable)Find(data interface{}) int{
	var colLId =0
	curpos := h.hashFunc(data,h.tableSize)//计算哈希值
	if h.theCells[curpos].kind!=Empty && h.theCells[curpos].data!=data {
		colLId +=1
		curpos :=2*curpos-1 //平方探测

		//如果越界
		if curpos>h.tableSize {
           curpos -=h.tableSize
		}
	}
	return curpos
}
func (h *HashTable) Insert (data interface{}){
	pos := h.Find(data) //查找位置
	entry := h.theCells[pos]
	//是否已经存在数据
	if entry.kind != legimate{
		entry.kind = legimate
		entry.data = data //插入数据
	}
}

func (h *HashTable) Empty(){
	for i:=0;i<h.tableSize;i++ {
		if h.theCells[i]==nil {
			continue
		}
		h.theCells[i].kind = Deleted //标志已删除
	}
}


func (h *HashTable) GetValue(index int)interface{}{
	if index>h.tableSize {
		return nil
	}
	entry := h.theCells[index]
	if entry.kind == legimate {
		return entry.data
	}else {
		return nil
	}
}