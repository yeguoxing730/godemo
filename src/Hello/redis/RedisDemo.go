package main

import "fmt"
import "github.com/garyburd/redigo/redis"

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("conn redis failed, err:", err)
		return
	}
	fmt.Println("connect redis success")
	//get set
	_, err = c.Do("Set", "name", "nick")
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.String(c.Do("Get", "name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
	//MSET MGET
	_, err = c.Do("MSet", "name", "nick", "age", "18")
	if err != nil {
		fmt.Println("MSet error: ", err)
		return
	}

	r2, err := redis.Strings(c.Do("MGet", "name", "age"))
	if err != nil {
		fmt.Println("MGet error: ", err)
		return
	}
	fmt.Println(r2)
	//hset hget
	_, err = c.Do("HSet", "names", "nick", "suoning")
	if err != nil {
		fmt.Println("hset error: ", err)
		return
	}

	r, err = redis.String(c.Do("HGet", "names", "nick"))
	if err != nil {
		fmt.Println("hget error: ", err)
		return
	}
	fmt.Println(r)
	//lpush lpop llen
	// 队列
	_, err = c.Do("lpush", "Queue", "nick", "dawn", 9)
	if err != nil {
		fmt.Println("lpush error: ", err)
		return
	}
	for {
		r, err = redis.String(c.Do("lpop", "Queue"))
		if err != nil {
			fmt.Println("lpop error: ", err)
			break
		}
		fmt.Println(r)
	}
	r3, err := redis.Int(c.Do("llen", "Queue"))
	if err != nil {
		fmt.Println("llen error: ", err)
		return
	}
	fmt.Println(r3)
	//expire
	_, err = c.Do("expire", "names", 5)
	if err != nil {
		fmt.Println("expire error: ", err)
		return
	}
	defer c.Close()

}