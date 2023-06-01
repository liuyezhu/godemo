package main

import (
	"fmt"
)

/*

   1）接口interface是一个自定义类型，接口类型具体描述了一系列方法的集合。

   2）接口类型是一种抽象的类型,它不会暴露出它代表的对象的内部值的结构和这个对象支持的基础操作的结合，

   他们只会展示出他们自己的方法。

   因此接口类型不能将其实例化。

   3）Go通过接口实现了鸭子类型(duck-typing)

*/

type Humaner interface {
	SayHi()

	//1)一般以er结尾

	//2)接口只有方法声明,没有实现,没有数据字段

	//3)接口可以匿名嵌入其他接口，或嵌入到结构中

}

type Student struct {
	name string

	score float64
}

//Student实现SayHi()方法

func (s *Student) SayHi() {

	fmt.Printf("Student[%s,%f] say hi!!!\n", s.name, s.score)

}

type Teacher struct {
	name string

	group string
}

func (t *Teacher) SayHia() {

	fmt.Printf("Teacher[%s,%s] say hi!!!\n", t.name, t.group)

}

func (t *Teacher) SayHi() {

	fmt.Printf("Teacher[%s,%s] say hi!!!\n", t.name, t.group)

}

type MyStr string

func (str MyStr) SayHi() {

	fmt.Printf("MyStr[%s] say hi!", str)

}

func WhoSayHi(i Humaner) {

	i.SayHi()

}

func main() {

	//接口的实现：1）接口是用来定义行为的类型。

	//2）这些被定义的行为不由接口直接实现,而是通过方法由用户定义的类型实现。

	//3）一个实现了这些方法的具体类型是这个接口类型的实例。

	//4)如果用户定义的类型实现了某个接口类型声明的一组方法，那么这个用户定义的类型的值就可以赋给这个接口类型的值。

	//这个赋值会把用户定义的类型的值存入接口类型的值。

	s := &Student{"ck_god", 88.88}

	t := &Teacher{"god_girl", "computer_programmer"}

	var tmp MyStr = "字符对象"

	s.SayHi()

	t.SayHi()

	tmp.SayHi()

	fmt.Println("\n==============\n")

	//多态--鸭子模型,调用同一接口，不同表现

	WhoSayHi(s)

	WhoSayHi(t)

	WhoSayHi(tmp)

	fmt.Println("\n==============\n")

	x := make([]Humaner, 3)

	x[0], x[1], x[2] = s, t, tmp

	for _, value := range x {

		value.SayHi()

	}

	fmt.Println("\n==============\n")

}
