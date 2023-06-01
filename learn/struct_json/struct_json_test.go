package struct_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Student struct {
	StudentId string `json:"student_id"`
	StudentName string `json:"student_name"`
	StudentClass string `json:"student_class"`
	StudentTeacher string `json:"student_teacher"`
}

func TestStudent(t *testing.T) {
	s :=&Student{StudentId: "1",StudentName: "wo",StudentClass: "math",StudentTeacher: "mm"}
	jsonString,err :=json.Marshal(s)
	fmt.Println(string(jsonString))
	fmt.Println(err)
}