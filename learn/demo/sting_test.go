package sel_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T){
	s :="A,B,C"
	parts := strings.Split(s,",")
	for _,part :=range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts,"--"))
}

func TestConv(t *testing.T){

	str := "100.00"


	strInt,err := strconv.Atoi(str)
	fmt.Println(err)
	fmt.Println(strInt,err)
//	s := strconv.Itoa(10)
//	t.Log("str"+s)
//t.Log(strconv.Atoi("10"))
//	//t.Log(10+strconv.Atoi("10"))
//
//
//	if i,err :=strconv.Atoi("10"); err ==nil {
//		t.Log(10+i)
//
//	}
//	if d, a :=strconv.Atoi("10"); a== nil {
//		t.Log(d)
//	}

}