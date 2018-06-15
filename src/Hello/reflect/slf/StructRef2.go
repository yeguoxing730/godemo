package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetAget() int {
	return p.Age
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func (p *Person) SetAget(name string) {
}

func main() {
	p := Person{
		Name: "Jim",
		Age:  27,
	}

	v1 := reflect.ValueOf(p)
	v2 := reflect.ValueOf(&p)

	fmt.Println("reflect.ValueOf(p)", v1)
	fmt.Println("reflect.ValueOf(&p)", v2)
	fmt.Println("Person Value Kind", v1.Kind())
	fmt.Println("&Person Value Kind", v2.Kind())
	fmt.Println("Person Value Type", v1.Type())
	fmt.Println("&Person Value Type", v2.Type())

	fmt.Println("reflect.ValueOf(p).Interface()", v1.Interface())
	fmt.Println("reflect.ValueOf(&p).Interface()", v2.Interface())

	fmt.Printf("\n======\n\n")

	//结构体是值传递，不可CanSet
	fmt.Println("reflect.ValueOf(p) CanSet: ", v1.CanSet())
	//结构体指针也是值传递，不可CanSet
	fmt.Println("reflect.ValueOf(&p) CanSet: ", v2.CanSet())
	//将结构体指针指向的结构出来，可CanSet
	fmt.Println("reflect.ValueOf(&p).Elem() CanSet: ", v2.Elem().CanSet())

	//可以直接修改CanSet的Object
	fmt.Println("Before Set: ", p.GetName())
	v2.Elem().FieldByName("Name").SetString("Kate")
	fmt.Println("After Set: ", p.GetName())

	fmt.Printf("\n======\n\n")

	//调用函数
	fmt.Println("Call GetName Method: ", v2.MethodByName("GetName").Call([]reflect.Value{}))
	fmt.Println("Call SetName Method: ", v2.MethodByName("SetName").Call([]reflect.Value{reflect.ValueOf("Tom")}))
	fmt.Println("Call GetName Method: ", v2.MethodByName("GetName").Call([]reflect.Value{}))
}