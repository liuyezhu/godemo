package sel

import (
	"fmt"
	"testing"
)

type a struct {

}

func DoSomething(p interface{})  {

	switch v:=p.(type) {
	case int:
		fmt.Println("Integer",v)
	case string:
		fmt.Println("String",v)
	default:
		fmt.Println("Unknow Type")
	}

	if i,ok := p.(int); ok {
		fmt.Println("Integer",i)
		return
	}
	if s,ok := p.(string);ok{
		fmt.Println("String",s)
		return
	}

	fmt.Println("Unknow Type")

}

func TestEmpty(t *testing.T){
	DoSomething(10)
	DoSomething("10")
	DoSomething(a{})
}

