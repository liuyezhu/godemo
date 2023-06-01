package helper

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

func ValidatorStruct(v interface{}) error {
	value := reflect.ValueOf(v)
	if value.Kind() == reflect.Ptr { // 如果传递进来的是指针
		value = value.Elem() // 拿到这个指针指向的地址
	}
	typ := value.Type() // 等价于reflect.TypeOf(v)

	for i := 0; i < value.NumField(); i++ { // 循环遍历属性
		tag := typ.Field(i).Tag.Get("binding") // 获得tag并且进行简单处理
		var (
			tagK string
			tagV string
		)
		equalIndex := strings.Index(tag, "=")
		if equalIndex != -1 {
			tagK = tag[0:equalIndex]
			tagV = tag[equalIndex+1:]
		}
		field := value.Field(i)
		switch field.Kind() { // 根据struct的属性进行判断，下面实现了两个简单的校验
		case reflect.String:
			if tag == "required" {
				if len(field.Interface().(string)) < 1 {
					return errors.New(tag + "not match" + typ.Field(i).Name)
				}
			}
		case reflect.Uint8:
			if tagK == "gt" {
				target, _ := strconv.Atoi(tagV)
				if field.Uint() <= uint64(target) {
					return errors.New(tag + "not match" + typ.Field(i).Name)
				}
			}
		}
	}
	return nil
}
