package sel_test

import (
	"fmt"
	"testing"
	"time"
)
//go协程机制
func TestGroutine(t *testing.T)  {
	for i :=0; i<10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Millisecond *50 )
}