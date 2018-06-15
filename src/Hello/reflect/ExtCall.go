package main

import (
	"fmt"
	"reflect"
)

func Halou(){
	fmt.Println("This is Halou函数！ 6666")
}

func main() {
	// Call()扩展
	h := Halou
	hVal := reflect.ValueOf(h)
	fmt.Println("hVal is reflect.Func ?", hVal.Kind() == reflect.Func)
	hVal.Call(nil)
}
