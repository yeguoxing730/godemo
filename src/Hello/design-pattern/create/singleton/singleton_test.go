package singleton

import (
	"testing"
	"fmt"
)

func TestSingleton(t *testing.T)  {
	user1 := GetUserInstance()
	user2 := GetUserInstance()
	if user1 != user2{
		t.Fatal("instance is not equal")
	}else {
		fmt.Println("equal user")
	}
}