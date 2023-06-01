package sel_test

import (
	"fmt"
	"testing"
	"time"
)


func service() string {
	time.Sleep(time.Millisecond*1001)
	return "Done"
}

func otherTask()  {
	fmt.Println("Working on something else")
	time.Sleep(time.Millisecond*100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T)  {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string{
	retCh := make(chan string,1)
	go func() {
		ret := service()
		fmt.Println("returned result")
		retCh <- ret
		fmt.Println("service exited")
	}()
	return  retCh
}

func TestAsyncService(t *testing.T)  {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Second*1)
}

//func TestSelect(t *testing.T)  {
//	select {
//	case ret := <-AsyncService():
//		t.Logf(ret)
//	case <-time.After(time.Millisecond*100):
//		t.Error("time out")
//
//	}
//}

//多路选择和超时
func TestS(t *testing.T) {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Println("received", i1, " from c1")
	case c2 <- i2:
		fmt.Println("sent ", i2, " to c2")
	case i3, ok := (<-c3):  // same as: i3, ok := <-c3
		if ok {
			fmt.Println("received ", i3, " from c3")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}


