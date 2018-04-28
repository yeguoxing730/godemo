package memcached

import (
	"github.com/pangudashu/memcache"
	"fmt"
)

func init()  {
	//server配置
	s1 := &memcache.Server{Address: "127.0.0.1:12000", Weight: 50}
	s2 := &memcache.Server{Address: "127.0.0.1:12001", Weight: 20}
	s3 := &memcache.Server{Address: "127.0.0.1:12002", Weight: 20}
	s4 := &memcache.Server{Address: "127.0.0.1:12003", Weight: 10}

	//初始化连接池
	mc, err := memcache.NewMemcache([]*memcache.Server{s1, s2, s3, s4})
	if err != nil {
		fmt.Println(err)
		return
	}

	//设置是否自动剔除无法连接的server，默认不开启(建议开启)
	//如果开启此选项被踢除的server如果恢复正常将会再次被加入server列表
	mc.SetRemoveBadServer(true)

	mc.Set("test_key",true)
	fmt.Println(mc.Get("test_key"))

	//...

	mc.Close()
}