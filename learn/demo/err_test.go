package sel_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

var LessThanTwoError = errors.New("n should be not less than 2")
var largerThanTwoError = errors.New("n should be not larger than 100")


func  getFibonacci(n int)  ([]int ,error){
	//if n<2 || n> 100 {
	//	return nil,errors.New("n should be in [2,100]")
	//}

	if n<2 {
		return nil, LessThanTwoError
	}

	if n>4 {
		return nil, largerThanTwoError
	}


	fibList := []int {1,2}
	for i := 2; i< n; i++ {
		fibList =append(fibList,fibList[i-2]+fibList[i-1])
	}
	return fibList ,nil

}

func TestFibonacci(t *testing.T)  {
	GetFibonacci2("1212")


	//if v, err := getFibonacci(5); err == nil {
	//	t.Log(v)
	//} else {
	//	if err == LessThanTwoError {
	//		//todo 你的错误操作
	//		fmt.Println("It is less.")
	//	}
	//	if err == largerThanTwoError {
	//		//todo 你的错误操作
	//		fmt.Println("It is larger.")
	//	}
	//}

}

func GetFibonacci2(str string)  {
	var (
		i int
		err error
		list []int
	)
	if i ,err =strconv.Atoi(str); err != nil {
		fmt.Println("Error\n",err)
		return
	}
	if list ,err = getFibonacci(i);err != nil{
		fmt.Println("Error\n",err)
		return
	}
	fmt.Println(list)
}