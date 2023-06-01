package sel

import (
	"fmt"
	"testing"
	"unsafe"
)


type Employee struct {
	Id string
	Name string
	Age int
}

func TestEmployeeObj(t *testing.T)  {
	e := Employee{"0","bob",20}
	e1 := Employee{Name: "M",Age:30}
	e2 := new(Employee)
	e2.Id = "2"
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T",&e)
	t.Logf("e is %T",e)
}

//结构体的实例会被复制地址会被变化
//func (e Employee) String() string {
//
//	fmt.Printf("Address is %x\n",unsafe.Pointer(&e.Name))
//
//	return fmt.Sprintf("ID:%s-Name:%s-Age:%d",e.Id,e.Name,e.Age)
//}



//没有数据的复制
func (e *Employee) String() string {
	fmt.Printf("Address is %x %d",unsafe.Pointer(&e.Name),1)
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d",e.Id,e.Name,e.Age)
}

func TestStructOperations(t *testing.T)  {
	e := Employee{"0","Bob",20}

	fmt.Printf("Address is %x\n",unsafe.Pointer(&e.Name))

	t.Log(e.String())
}

