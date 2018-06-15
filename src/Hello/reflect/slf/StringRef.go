package main

import (
	"fmt"
	"reflect"
)

func main()  {
	a:="reflect string"
	aT := reflect.TypeOf(&a)
	aV := reflect.ValueOf(&a)
	fmt.Println(aT.Kind().String())
	fmt.Println(aV.Elem().Interface())
	s := reflect.New(aT).Interface()
	fmt.Println(&s)
}
