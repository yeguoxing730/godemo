package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Response struct {
	Length int         `json:"length"`
	Data   interface{} `json:"data"`
}

type Instance struct {
	Id      int    `json:"id"`
	OrderId int    `json:"order_id"`
	Name    string `json:"instance_name"`
}

type Order struct {
	Id   int    `json:"id"`
	Name string `json:"order_name"`
}

func MarshalData(data interface{}) []byte {
	var isSlice = false
	var resp *Response
	var ret []byte
	var err error

	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Slice {
		isSlice = true
	}
	if isSlice {
		resp = &Response{
			Length: v.Len(),
			Data:   data,
		}
		ret, err = json.Marshal(resp)
	} else {
		ret, err = json.Marshal(data)
	}

	if err != nil {
		return []byte("data error")
	} else {
		return ret
	}
}

func main() {
	orders := []*Order{
		&Order{
			Id:   1,
			Name: "first",
		},
		&Order{
			Id:   2,
			Name: "second",
		},
	}

	instances := []*Instance{
		&Instance{
			Id:      1,
			OrderId: 1,
			Name:    "first",
		},
		&Instance{
			Id:      2,
			OrderId: 2,
			Name:    "second",
		},
	}

	one := &Order{
		Id:   10,
		Name: "only",
	}

	fmt.Println(string(MarshalData(instances)))
	fmt.Println(string(MarshalData(orders)))
	fmt.Println(string(MarshalData(one)))
}