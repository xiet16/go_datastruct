package DistributionSort

import (
	"fmt"
	"sync"
	"time"
)

type SyncMap struct {
	mymap map[string]string
	*sync.RWMutex
}

var smap SyncMap
var done chan bool //通道，是否完成

func Write()  {
	keys :=[]string{"1","2","3"}
	for _,k :=range keys{
		smap.Lock()
		smap.mymap[k] =k
		smap.Unlock()
		time.Sleep(1*time.Second)
	}
	done <-true
}

func Write2()  {
	keys :=[]string{"1","2","3"}
	for _,k :=range keys{
		smap.Lock()
		smap.mymap[k] =k
		smap.Unlock()
		time.Sleep(1*time.Second)
	}
	done <-true
}

func Read()  {
	smap.RLock()
	fmt.Println("readlock")
	for k,v :=range smap.mymap{
		fmt.Println(k,"-->",v)
	}
	smap.RUnlock()
}