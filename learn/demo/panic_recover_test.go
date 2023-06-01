package sel

import (
	"errors"
	"fmt"
	"runtime"
	"testing"
)

func TestPanicVxExit(t *testing.T)  {

	defer func() {
		if err := recover() ; err !=nil {
			fmt.Println("recovered from",err)
		}
	}()
	fmt.Println(runtime.Version())

	fmt.Println("Start")
	panic(errors.New("Something wrong!"))

	//os.Exit(-1)
}

