package channel1

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

//仅需任务完成
func AllResponse() string{
	numOfRunner := 10
	ch :=make(chan string,numOfRunner)
	for i :=0;i<numOfRunner;i++ {
		go func(i int) {
			ret := runTask(i)
			ch <-ret
		}(i)
	}
	finalRet :=""
	for j:=0;j<numOfRunner;j++{
		finalRet += <-ch +"\n"
	}
	return finalRet
}

func runTask(id int) string  {
	time.Sleep(10*time.Millisecond)
	return fmt.Sprintf("The result is from %d",id)
}

func TestAllResponse(t *testing.T)  {
	t.Log("Before",runtime.NumGoroutine())
	t.Log(AllResponse())
	t.Log("After:",runtime.NumGoroutine())
}