package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	UserName string
	Age      int
}

func (user *Student) Sing() {
	fmt.Println(user.UserName + " is singing.....")
}
func (user *Student) ChangeAge(newAge int) {
   user.Age = newAge
}

func main() {
	stu := Student{"lilan", 23}
	stuObj := reflect.ValueOf(&stu).Elem()
	stuInf := reflect.TypeOf(&stu)
	stuType := stuObj.Type()
	fmt.Println(stuType.Name())
	fields := stuObj.NumField()
	for i := 0; i < fields; i++ {
		typeName := stuObj.Field(i).Type().Name()
		fieldVal := stuObj.Field(i).Interface()
		fieldName := stuType.Field(i).Name
		fmt.Printf("%d: %s %s = %v\n", i, fieldName, typeName, fieldVal)
	}
	stuObj.Field(0).SetString("new-user-name")
	stuObj.Field(1).SetInt(32)
	fmt.Println(stu)

	methods := stuObj.NumMethod()
	fmt.Println(methods)
	methods = stuType.NumMethod()
	fmt.Println(methods)
	methods = stuInf.NumMethod()
	fmt.Println(methods)

	for i := 0; i < methods; i++ {
		methodName := stuInf.Method(i).Name
		methodDefType := stuInf.Method(i).Func.Type()
		fmt.Println(methodName)
		fmt.Println(methodDefType)
	}
	m := stuInf.Method(1)
	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(&stu)
	m.Func.Call(params)

	m2 := stuInf.Method(0)
	params2 := make([]reflect.Value, 2)
	params2[0] = reflect.ValueOf(&stu)
	params2[1] = reflect.ValueOf(44)
	m2.Func.Call(params2)
	fmt.Println(stu)
}
