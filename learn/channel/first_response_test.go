package channel1_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

//仅需任务完成
func FirstResponse() string{
	numOfRunner := 10
	ch :=make(chan string,numOfRunner)
	for i :=0;i<numOfRunner;i++ {
		go func(i int) {
			ret := runTask(i)
			ch <-ret
		}(i)
	}
	return <-ch
}

func runTask(id int) string  {
	time.Sleep(10*time.Millisecond)
	return fmt.Sprintf("The result is from %d",id)
}

func TestFirstResponse(t *testing.T)  {
	t.Log("Before",runtime.NumGoroutine())
	t.Log(FirstResponse())
	t.Log("After:",runtime.NumGoroutine())
}