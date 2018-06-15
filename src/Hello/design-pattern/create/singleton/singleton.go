package singleton

import "sync"

type User struct{
	UserName string
	Age int
}

var user *User
var once sync.Once
func GetUserInstance() *User{
	once.Do(func() {
		user = &User{"ye",23}
	})
	return user
}
///使用懒惰模式的单例模式，使用双重检查加锁保证线程安全