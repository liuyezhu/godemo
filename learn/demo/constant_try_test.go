package sel_test

import (
	"fmt"
	"testing"
)

const (
	Monday =1 + iota
	Tuesday
	Wednesday
)
const (
	Readable = 1 << iota
	Writable
	Executable
	test
)

func TestConstantTry(t *testing.T)  {
	fmt.Println(7 & 2)
	//t.Log(Monday)
}

func TestCon(t *testing.T)  {
	//a :=7
	//0111
	a :=10
	//1010

	t.Log(Readable)      //1  0001
	t.Log(a& Readable)   //0
	t.Log(Writable)      //2  0010
	t.Log(a& Writable)   //2  0010
	t.Log(Executable)    //4   0101
	t.Log(a& Executable) //0   0

	//t.Log(test)
	t.Log(a&Readable == Readable,a&Writable == Writable,a&Executable == Executable)
}