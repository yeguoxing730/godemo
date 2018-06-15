package proxy

import "testing"

func TestProxy_Do(t *testing.T) {
	sub := RealSubject{}
	subject := NewProxySubject(sub)
	subject.Do()
}