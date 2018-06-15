package simple_factory

import "fmt"

type Product interface {
	GetName() string
}
type ProductBase struct {
	Name string
}

type Mice struct {
	*ProductBase
}

func (mice *Mice) GetName()string {
	fmt.Println("mice name")
	return "mice name"
}

type Tomato struct {
	*ProductBase
}

func (tomata *Tomato)GetName()string  {
	fmt.Println("tomato name")
	return "tomato name"
}

func NewProduct(ty int)Product {
	if ty ==1{
		return &Mice{}
	}
	if ty ==2{
		return &Tomato{}
	}
	return nil
}

//go 语言没有构造函数一说，所以一般会定义NewXXX函数来初始化相关类。
//NewXXX 函数返回接口时就是简单工厂模式，也就是说Golang的一般推荐做法就是简单工厂。