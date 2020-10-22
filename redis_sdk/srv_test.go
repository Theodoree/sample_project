package redis_sdk

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
    "net"
    "sync"
    "testing"
    "time"
)

func TestNewSrv(t *testing.T) {

    s := redis.NewPool(func() (redis.Conn, error) {
        conn, err := net.Dial("tcp", "localhost:6379")
        if err != nil {
            return nil, err
        }

        redisConn := redis.NewConn(conn, time.Second*5, time.Second*5)

        //redisConn.Do("AUTH")

        redisConn.Do("PING")
        return redisConn, nil

    }, 1024)

    s.TestOnBorrow = func(c redis.Conn, t time.Time) error {
        _, err := c.Do("PING")
        if err != nil {
            fmt.Println(err)
        }
        return err
    }

    srv := NewSrv(s)

    var wait sync.WaitGroup
    wait.Add(20)

    cb := func() {
        wait.Done()
    }

    for i := 0; i < 20; i++ {
        go threadLock(&srv, "礼包A", i, cb)
    }



    wait.Wait()

}

func threadLock(s *Srv, key string, GoroutineNum int, cb func()) {

    defer cb()
    for {
        ok, err := s.GetLock(key)
        if err != nil {
            fmt.Printf("%v \n", err)
            continue
        }
        // 获取到锁
        if ok {

            if s.RetryLock(key, time.Now().Add(time.Second*10).Unix()) {
                s.UnLock(key)
                println(GoroutineNum, "获取到锁了")
                return
            }

        } else {
            ok, err := s.Lock(key, time.Now().Add(time.Second*10).Unix())
            if err != nil {
                return
            }

            // 睡50 ms
            // 如果锁上直接返回
            if ok {
                s.UnLock(key)
                println(GoroutineNum, "获取到锁了")
                return
            }
        }

    }

}
