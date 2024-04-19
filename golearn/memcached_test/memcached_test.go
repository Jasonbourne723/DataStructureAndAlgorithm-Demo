package memcachedtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gomodule/redigo/redis"
)

func BenchmarkMemcached(b *testing.B) {

	mc := memcache.New("127.0.0.1:11211")

	for i := 0; i < b.N; i++ {
		key := strconv.FormatInt(int64(i), 10)
		mc.Set(&memcache.Item{
			Key:   key,
			Value: []byte(key),
		})
	}

	for i := 0; i < b.N; i++ {
		key := strconv.FormatInt(int64(i), 10)
		_, _ = mc.Get(key)
	}

}

func BenchmarkRedis(b *testing.B) {

	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	defer c.Close()
	for i := 0; i < b.N; i++ {
		key := strconv.FormatInt(int64(i), 10)

		_, _ = c.Do("Set", key, []byte(key))

	}

	for i := 0; i < b.N; i++ {
		key := strconv.FormatInt(int64(i), 10)
		_, _ = c.Do("Get", key)
	}
}
