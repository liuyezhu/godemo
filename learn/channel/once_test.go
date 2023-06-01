package channel1

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
	Name string `json:"name"`
	Age int
	High bool
	Sex string
}

var singleInstance  *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Obj")
		singleInstance = new(Singleton)
		singleInstance.Name ="sdsd"
		singleInstance.Age =123
		singleInstance.High =true
		singleInstance.Sex ="sd"
	})
	return singleInstance
}

func TestGetSingletonObj(t *testing.T){
	var wg sync.WaitGroup
	for i:=0;i<10;i++ {
		wg.Add(1)
		go func() {
			obj :=GetSingletonObj()
			fmt.Println(obj)
			fmt.Printf("%x\n",unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
