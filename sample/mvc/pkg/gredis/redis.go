package gredis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/Theodoree/sample_project/sample/mvc/pkg/setting"
	"time"
	"github.com/gin-gonic/gin/json"
)

var RedisConn *redis.Pool

func Setup() {
	RedisConn = &redis.Pool{
		MaxActive: setting.RedisSetting.MaxActive,
		MaxIdle:   setting.RedisSetting.MaxIdle,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(`tcp`, setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
					return nil, err
				}
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

}

func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value) //添加
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time) //设置超时时候
	if err != nil {
		return err
	}

	return nil
}

func Exists(key string) bool { //是否存在
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func Get(key string) ([]byte, error) { //Get
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func Delete(key string) (bool, error) { //DELETE
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(key string) error { //模糊删除
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}