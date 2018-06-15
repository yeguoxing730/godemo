package proxy

import (
	"fmt"
)

type Subject interface {
	Do() string
}

type RealSubject struct {}

func (RealSubject)Do()string  {
	fmt.Println("do real subject")
	return "real"
}

type Proxy struct {
	realObj Subject
}

func NewProxySubject(subject Subject)Subject {
	proxy := Proxy{}
	proxy.realObj = subject
	return  proxy
}
func (p Proxy)Do() string {
	fmt.Println("pre execute")
	out := p.realObj.Do()
	fmt.Println("after execute")
	return out
}