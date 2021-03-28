package Hash

import "fmt"

func HashTest()  {
	fmt.Println("mysha :",MySha("xiet",100))
	fmt.Println("mysha:",MySha("xiet16",100))

	fmt.Println("mysha :",MySha256("xiet",100))
	fmt.Println("mysha:",MySha256("xiet16",100))
}

func HashTableTest()  {
	table,err:= NewHashTable(100,MySha256)

	if err!=nil {
		fmt.Println("create table error :",err.Error())
		return
	}

	table.Insert("xiet")
	table.Insert("daibo")
	table.Insert("xiaopeng")
	index:=table.Find("xiet")
    fmt.Println(index ,"-->",table.GetValue(index))
}