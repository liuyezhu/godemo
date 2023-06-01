package channel1_test

import (
	"context"
	"fmt"
	"testing"
)

//Content任务取消
func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:

		return false
	}
}


func TestCancelChannel(t *testing.T)  {
	ctx,cancel :=context.WithCancel(context.Background())
	for i:=0;i<5;i++ {
		go func(i int, ctx context.Context){
			for  {
				if isCancelled(ctx) {
					break
				}
				//time.Sleep(time.Millisecond*1)
			}
			fmt.Println(i,"cancel")
		}(i,ctx)
	}
	cancel()
	//time.Sleep(time.Second*1)
}