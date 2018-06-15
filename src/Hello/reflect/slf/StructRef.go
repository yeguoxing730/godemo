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

	t1 := reflect.TypeOf(p)
	t2 := reflect.TypeOf(&p)

	fmt.Println("Person Type", t1)
	fmt.Println("&Person Type", t2)

	//结构体类型Kind
	fmt.Println("Person Kind:", t1.Kind())
	//结构体指针类型Kind
	fmt.Println("&Person Kind:", t2.Kind())

	//结构体名称
	fmt.Println("Person Name:", t1.Name())
	//指针为空
	fmt.Println("Person Name:", t2.Name())

	fmt.Printf("\n======\n\n")

	//只有结构体才有Field，假如t2.NumField()会报错
	fmt.Println("Person Num Fields", t1.NumField())
	for i := 0; i < t1.NumField(); i++ {
		//获取各个域的名称，类型，json tag
		fmt.Printf("Person Field %d: [Name: %s, Type: %s, Tag: %s]\n", i, t1.Field(i).Name, t1.Field(i).Type, t1.Field(i).Tag.Get("json"))
	}

	fmt.Printf("\n======\n\n")

	// 上面四个方法均为结构体指针的方法，而不是结构体的方法，t1.NumMethod会报错
	fmt.Println("Person Num Methods", t2.NumMethod())
	for i := 0; i < t2.NumMethod(); i++ {
		//方法的名称，类型
		fmt.Printf("Person Method %d: [Name: %s, Type: %s]\n", i, t2.Method(i).Name, t2.Method(i).Type)
		//方法的输入参数，其中第一个参数为结构体指针本身
		fmt.Printf("  NumIn: %d\n", t2.Method(i).Type.NumIn())
		for j := 0; j < t2.Method(i).Type.NumIn(); j++ {
			param := t2.Method(i).Type.In(j)
			fmt.Printf("    param %d: [kind: %s]\n", j, param.Kind())
		}

		//方法的输出参数
		fmt.Printf("  NumOut: %d\n", t2.Method(i).Type.NumOut())
		for j := 0; j < t2.Method(i).Type.NumOut(); j++ {
			param := t2.Method(i).Type.Out(j)
			fmt.Printf("    param %d: [kind: %s]\n", j, param.Kind())
		}
	}
}
