package sel_test

import (
	"fmt"
	"testing"
)


type Code string
type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {

}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}
type JavaProgrammer struct {

}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World!\")"
}


func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}


func TestProgrammer(t *testing.T)  {
	goProg := &GoProgrammer{}
	javaProg := &JavaProgrammer{}
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}


type Age int32
func (age Age) InitAge() int32 {
	if age == 0 {
		age = 10
		return int32(age)
	}
	return int32(age)
}

func TestAction(t *testing.T) {
	age := Age(0)
	initAge := age.InitAge()
	fmt.Println(initAge)
}
