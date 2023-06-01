package sel_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int,int)  {
	return rand.Intn(10),rand.Intn(20)
}




//分开写参数类型
func add(x int, y int) int {
	return  x +y
}
//相同参数类型
func add1 (x,y int) int {

	return  x +y
}
//多返回值
func addAndSub(x int,y int) (int,int){
	var add = x +y
	var sub = x -y
	return add,sub
}



func timeSpent(innerA func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		fmt.Println(start)
		ret := innerA(n)
		fmt.Println("time spent:",time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int)int  {
	time.Sleep(time.Second*1)
	return op
}

func slowF(a int) int  {
	return a
}
//命名返回
func SList(num int) (dig2 , dig1 int)  {
	dig1 = num *4
	dig2 = num *9
	return  dig2,dig1
}
func TestFn(t *testing.T)  {
	a,_ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFun)
	ta := timeSpent(slowF)

	t.Log(tsSF(10))
	t.Log(ta(10))
	t.Log(	add(1,2))
	t.Log(	add1(1,2))
	t.Log(	addAndSub(1,2))
	_,c := addAndSub(1,2)
	t.Log(c)
	t.Log(SList(1))

	t.Log(Sum(1,1,2,1))
	t.Log(Sum(1,1,2,11111))
}
//求和
func Sum(ops ...int) int  {
	ret :=0
	for _,op :=range ops{
		ret += op
	}
	return  ret
}

func Clear (){

	fmt.Println("clear resources")
}


//defer语句会延迟函数的执行上层函数返回
//延迟调用的参数会立刻生成，但是在上层函数都不会被调用
//延迟的函数调用会被压入一个栈中。但函数返回时，会按照后进先出的顺序调用的函数调用
func TestDefer(t *testing.T)  {
	//相当与PHP try{}catch{}
	defer Clear() //用于释放锁等操作
	fmt.Println("start")
	panic("err")
}
