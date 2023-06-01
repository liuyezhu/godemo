package sel

import (
	"fmt"
	"testing"
)

type Pet struct {
	
}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println("",host)
}

//type Dog struct_json {
//	p *Pet
//}

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	fmt.Println("Wang!")
}
//
//func (d *Dog) SpeakTo(host string) {
//	d.Speak()
//	fmt.Println(host)
//}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.Speak()
	//dog.SpeakTo("chao")
}