package main

import (
	"reflect"
	"fmt"
)

func Halou2(s string) string {
	return "接受到参数：" + s + ", 但这是返回值！"
}

func main() {
	h2 := reflect.ValueOf(Halou2)
	params := []reflect.Value{
		reflect.ValueOf("参数1"),
	}
	h2ReturnVal := h2.Call(params)
	fmt.Println(h2ReturnVal)
}