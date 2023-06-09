package channel1

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup)  {
	go func() {
		for i:=0;i<10;i++ {
			ch<-i
		}
		close(ch)
		//ch<-11
		wg.Done()

	}()
}

func clone(ch chan int) {
	
}

//func clone(ch chan int) {
//
//}

//func clone(ch chan int) {
//
//}

//channel的关闭和广播
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i:=0;i<10;i++{
			if data,ok:=<-ch;ok{
				fmt.Println(data)
			}else {

				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T)  {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch,&wg)
	wg.Add(1)
	dataReceiver(ch,&wg)
	wg.Add(1)
	dataReceiver(ch,&wg)
	wg.Wait()
}