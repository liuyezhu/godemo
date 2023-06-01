package sel

import (
	"testing"
)

func TestMapWithFunValue(t *testing.T){
	m := map[int]func(op int)int{}
	m[1] = func(op int) int {
		return  op
	}
	m[2] = func(op int) int {
		return op*op
	}
	m[3] = func(op int) int {
		return op*op*op
	}
	t.Log(m[1](2),m[2](2),m[3](3))
}

func TestMapForSet(t *testing.T){

	mySet := map[interface{}]interface{}{}
	mySet["sd"] = 1
	mySet[3]="1212"
	n := "sd"
	if mySet[n] == true {
		t.Logf("%s is existing",n)
	}else {
		t.Logf("%s is not existing",n)
	}
	mySet[3] = 2
	t.Log(mySet)
	t.Log(len(mySet))
	delete(mySet,1)
	if mySet[n] ==true{
		t.Logf("%s is existing",n)
	}else {
		t.Logf("%s is not existing",n)
	}

}