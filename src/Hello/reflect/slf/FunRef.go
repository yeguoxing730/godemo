package slf

import (
	"fmt"
	"strconv"
	"reflect"
)

func func1(){
	fmt.Println("func1 execute....")
}
func func2( in int)string{
	fmt.Println("func2 execute....")
	return  strconv.Itoa(in+100)
}

func makeFunc(fu interface{})interface{}{
	fuType := reflect.TypeOf(fu)
	if fuType.Kind() != reflect.Func{
		panic("not function type")
	}
	fuVal := reflect.ValueOf(fu)
	Fun := reflect.MakeFunc(fuType, func(args []reflect.Value) (results []reflect.Value) {
		fmt.Println("before execute")
		out := fuVal.Call(args)
		fmt.Println("after execute")
		return  out
	})
	return Fun.Interface()
}
func main()  {
	f1 := makeFunc(func1).(func())
	f11 := makeFunc(f1).(func())
	f11()
	f2 := makeFunc(func2).(func(int)string)
	res := f2(33)
	fmt.Println(res)
}