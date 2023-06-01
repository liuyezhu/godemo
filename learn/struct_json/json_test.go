package struct_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Stu struct {
	Name string `json:"name"`
	Age int
	High bool
	Sex string
	Class *Class `json:"class"`
}

type Class struct {
	Name string
	Grade int
}

func TestStu(t *testing.T) {
	stu :=Stu{
		Name: "liquefy",
		Age: 26,
		High: false,
		Sex: "male",
	}
	 c := new(Class)
	 c.Grade =211
	 c.Name ="sss"
	 stu.Class =c
	//Marshal失败时err!=nil
	jsonStu ,err :=json.Marshal(stu)
	if err !=nil {
		fmt.Println("生成json字符串错误")
	}
	fmt.Println(string(jsonStu))
}
