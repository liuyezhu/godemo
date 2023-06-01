package sel

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond*50)
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
//B 10 2 12
//BB 10 12 22

//
func calc(index string, a,b int) int {

	ret :=  a + b
	fmt.Println(index,a,b,ret)
	return ret
}

//func TestCalc(t *testing.T) {
//	x := 1
//	y := 2
//	defer calc("AA",x,calc("A",x,y))
//	x=10
//	defer calc("BB",x,calc("B",x,y))
//    y=20
//}