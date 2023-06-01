package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string `validate:"required" label:"名字"`
	Age   uint8  `validate:"gte=0,lte=130"`
	Email string `validate:"required,email"`
}

func main() {
	user := User{
		Name: "Helios",
		Age:  0,
	}
	fmt.Println(reflect.ValueOf(user), reflect.TypeOf(user)) // {Helios 0} main.User
	vv := reflect.ValueOf(user)
	fmt.Println(reflect.TypeOf(vv))
	vt := reflect.TypeOf(user)
	fmt.Println(vt)
	for i := 0; i < vv.NumField(); i++ {
		fieldValue := vv.Field(i)
		fieldTyp := vt.Field(i)
		fmt.Println(fieldValue) // Helios
		fmt.Println(fieldTyp)   //{Name  string validate:"required" 0 [0] false}
		name := fieldTyp.Tag
		fmt.Println(name)

	}
}
