package channel1

//任务取消
func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel1(cancelChan chan struct{})  {
	cancelChan <- struct{}{}
}

func cancel2(cancelChan chan struct{})  {
	close(cancelChan)
}

//func TestCancelChannel(t *testing.T)  {
//	cancelChan :=make(chan struct{},0)
//	for i:=0;i<5;i++ {
//		go func(i int, cancelCh chan struct{}){
//			for  {
//				if isCancelled(cancelCh) {
//					break
//				}
//				//time.Sleep(time.Millisecond*1)
//			}
//			fmt.Println(i,"cancel")
//		}(i,cancelChan)
//	}
//	cancel2(cancelChan)
//	time.Sleep(time.Second*1)
//}