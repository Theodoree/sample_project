package redis_sdk

import (
    "errors"
    "github.com/gomodule/redigo/redis"
)

type Srv struct {
    pool *redis.Pool
}

func NewSrv(pool *redis.Pool) Srv {
    return Srv{pool: pool}
}

func (s *Srv) GetLock(key string) (bool, error) {
    _, ok, err := s.Get(key)
    if err != nil {
        if errors.Is(err, redis.ErrNil) {
            return false, nil
        }
        return false, err
    }

    return ok, nil
}

func (s *Srv) TTL(key string) (int64, error) {
    if len(key) == 0 {
        return 0, nil
    }
    conn := s.pool.Get()
    defer conn.Close()

    ttl, err := s.ttl(conn, key)

    return ttl, err
}
func (s *Srv) ttl(conn redis.Conn, key string) (int64, error) {
    ttl, err := redis.Int64(conn.Do("TTL", key))
    if err != nil {
        return 0, err
    }
    return ttl, nil
}

func (s *Srv) SetNx(key string, val interface{}, ts int64) (ok bool, err error) {
    if len(key) == 0 {
        return
    }
    conn := s.pool.Get()
    defer conn.Close()

    ok, err = s.setNx(conn, key, val, ts)
    return ok, err
}
func (s *Srv) setNx(conn redis.Conn, key string, val interface{}, ts int64) (ok bool, err error) {
    var arr []int64
    _, err = conn.Do("MULTI")
    if err != nil {
        return
    }
    _, err = redis.String(conn.Do("SETNX", key, val))
    if err != nil {
        return
    }
    arr, err = redis.Int64s(conn.Do("EXEC"))
    if err != nil {
        return
    }

    // SETNX成功,然后再设置过期时间
    if len(arr) > 0 && arr[0] == 1 {
        ok = true
        _, err = conn.Do("EXPIREAT", key, ts)
        if err != nil {
            return
        }
    }

    return
}

func (s *Srv) Lock(key string, ts int64) (ok bool, err error) {
    conn := s.pool.Get()
    defer conn.Close()
    ok, err = s.lock(conn, key, ts)

    return
}
func (s *Srv) lock(conn redis.Conn, key string, ts int64) (ok bool, err error) {
    ok, err = s.setNx(conn, key, 0, ts)
    return
}

func (s *Srv) UnLock(key string) {
    s.Del(key)
}

func (s *Srv) RetryLock(key string, ts int64) bool {

    conn := s.pool.Get()
    defer conn.Close()

    val, err := s.ttl(conn, key)

    if err != nil {
        return false
    }
    if val <= 0 {
        ok, err := s.lock(conn, key, ts)
        if err != nil {
            return false
        }
        // 如果锁上直接返回
        if ok {
            return true
        }
    }
    return false
}

func (s *Srv) Del(key string) (err error) {
    if len(key) == 0 {
        return
    }
    conn := s.pool.Get()
    defer conn.Close()
    _, err = conn.Do("del", key)
    if err != nil {
        return
    }
    return
}

func (s *Srv) Set(key string, val interface{}, ts int64) (err error) {
    if len(key) == 0 {
        return
    }
    conn := s.pool.Get()
    defer conn.Close()

    _, err = conn.Do("MULTI")
    if err != nil {
        return
    }
    _, err = conn.Do("SET", key, val)
    if err != nil {
        return
    }
    _, err = conn.Do("EXPIREAT", key, ts)
    if err != nil {
        return
    }
    _, err = conn.Do("EXEC")
    if err != nil {
        return
    }
    return

}
func (s *Srv) Get(key string) ([]byte, bool, error) {
    if len(key) == 0 {
        return nil, false, nil
    }
    conn := s.pool.Get()
    defer conn.Close()
    buf, err := redis.Bytes(conn.Do("GET", key))
    if err != nil {
        return nil, false, err
    }
    return buf, true, nil
}

func (s *Srv) SADD(setKey string, value interface{}) (err error) {
    if len(setKey) == 0 {
        return
    }
    conn := s.pool.Get()
    defer conn.Close()

    _, err = conn.Do("MULTI")
    if err != nil {
        return
    }
    _, err = conn.Do("SADD", setKey, value)
    if err != nil {
        return
    }
    _, err = conn.Do("EXEC")
    if err != nil {
        return
    }

    return
}

func (s *Srv) SREM(setKey string, value string) (err error) {
    if len(setKey) == 0 {
        return
    }
    conn := s.pool.Get()
    defer conn.Close()

    _, err = conn.Do("MULTI")
    if err != nil {
        return
    }
    _, err = conn.Do("SREM", setKey, value)
    if err != nil {
        return
    }
    _, err = conn.Do("EXEC")
    if err != nil {
        return
    }

    return nil
}

func (s *Srv) SRANDMEMBER(setKey string, count uint) ([]interface{}, bool, error) {
    if len(setKey) == 0 {
        return nil, false, nil
    }
    conn := s.pool.Get()
    defer conn.Close()
    ret, err := redis.Values(conn.Do("SRANDMEMBER", setKey, count))
    if err != nil {
        return nil, false, err
    }
    return ret, true, nil
}
