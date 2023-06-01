package sel_test

//容量是否可伸缩
//是否可以比较
import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestArrayInit(t *testing.T)  {

	var arr [3]int
	arr1 := [4]int{1,2,3,4}
	arr3 := [...] string {"a","a","b"}
	arr1[1] =5
	t.Log(arr[1],arr[2])
	t.Log(arr1,arr3)

	var jsonBlob = []byte(`[
{"Name": "Platypus", "Order": "Monotremata"}
]`)
	type Animal struct {
		Name  string
		Order string
		ID    int
	}
	animals := &Animal{ID: -9999}
	err := json.Unmarshal(jsonBlob, animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}

func TestArrayTravel(t *testing.T)  {

	arr3 := [...]int{1,3,3,5}
	for i :=0;i<len(arr3);i++ {
		t.Log(arr3[i])
	}
	for idx,e := range arr3{
		t.Log(idx,e)
	}

}


func TestArrayType(t *testing.T) {

	arr3 :=[...]interface{}{1,"aa",true}
	for idx,e := range arr3{
		t.Log(idx,e)
	}
}

//数组切片
func TestArraySection(t *testing.T) {

	arr3 :=[...]interface{}{1,"aa",true,1.333}
	arr3Sec := arr3[:]
	arr3S := arr3[1:]

	t.Log(arr3Sec)
	t.Log(arr3S)
}
