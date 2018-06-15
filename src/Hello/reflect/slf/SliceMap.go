package main

import (
	"reflect"
	"fmt"
)

func main(){
	aMap := make(map[string]int)
	bSlice := make([]int,0)

	aMapT := reflect.TypeOf(aMap)
	bSliceT := reflect.TypeOf(bSlice)

	newMap := reflect.MakeMap(aMapT)
	key := reflect.ValueOf("key")
	val := reflect.ValueOf(32)
	newMap.SetMapIndex(key,val)
	rsMap := newMap.Interface().(map[string]int)
	fmt.Println(rsMap)


	newSlice := reflect.MakeSlice(bSliceT,0,0)
	newSlice = reflect.Append(newSlice,val)
	rsSlice := newSlice.Interface().([]int)
	fmt.Println(rsSlice)

}
