package main

import (
	"reflect"
	"fmt"
)

type Foo struct {
	A int `tag1:"Tag1" tag2:"Second Tag"`
	B string
}
func main() {
	// 反射的使用
	s := "String字符串"
	fo := Foo{A: 10, B: "字段String字符串"}

	sVal := reflect.ValueOf(s)
	// 在没有获取指针的前提下，我们只能读取变量的值。
	fmt.Println(sVal.Interface())

	sPtr := reflect.ValueOf(&s)
	sPtr.Elem().Set(reflect.ValueOf("修改值1"))
	sPtr.Elem().SetString("修改值2")
	// 修改指针指向的值，原变量改变
	fmt.Println(s)
	fmt.Println(sPtr) // 要注意这是一个指针变量，其值是一个指针地址

	foType := reflect.TypeOf(fo)
	foVal := reflect.New(foType)
	// foVal.Elem().Field(0).SetString("A") // 引发panic
	foVal.Elem().Field(0).SetInt(1)
	foVal.Elem().Field(1).SetString("B")
	f2 := foVal.Elem().Interface().(Foo)
	fmt.Printf("%+v, %d, %s\n", f2, f2.A, f2.B)
}
