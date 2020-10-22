package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

type accessController struct {
    m     map[string]*Call
    mutex sync.Mutex
    pool  sync.Pool
}

type Call struct {
    result interface{}
    err    error
    wait   sync.WaitGroup
    onCall int64
}

func (c *Call) reset() {
    c.result = nil
    c.err = nil
    c.onCall = 0
}
func (c *Call) addOnCall() {
    atomic.AddInt64(&c.onCall, 1)
}
func (c *Call) subOnCall() {
    atomic.AddInt64(&c.onCall, -1)
}
func (c *Call) notUse() bool {
    return atomic.LoadInt64(&c.onCall) == 0
}

var (
    controller = newAccessController()
)

func (c *accessController) getCall() *Call {
    return c.pool.Get().(*Call)
}

func (c *accessController) putCall(call *Call) {
    c.pool.Put(call)
}

func newAccessController() *accessController {
    return &accessController{
        pool: sync.Pool{
            New: func() interface{} {
                return newCall()
            },
        },
        mutex: sync.Mutex{},
        m: make(map[string]*Call),
    }
}

func newCall() *Call {
    return &Call{
        result: nil,
        err:    nil,
        wait:   sync.WaitGroup{},
    }

}

func Access(accessKey string, fn func() (interface{}, error)) (result interface{}, err error) {

    // 先上锁
    controller.mutex.Lock()
    if call, ok := controller.m[accessKey]; ok {
        // 增加引用量
        call.addOnCall()
        controller.mutex.Unlock()

        // 这里阻塞等待条件成立
        call.wait.Wait()

        // 设置返回值
        result = call.result
        err = call.err

        // 因为是atomic,所以不存在资源竞争的问题,这里不需要锁,减少引用量
        call.subOnCall()

        controller.mutex.Lock()
        if call.notUse() {
            // 如果没有人使用了,那么就将其放回到临时等待池
            controller.putCall(call)
        }
        controller.mutex.Unlock()

        return result, err
    }

    call := controller.getCall()
    call.wait.Add(1)
    controller.m[accessKey] = call
    controller.mutex.Unlock()

    call.result, call.err = fn()
    // 条件满足,上面放行
    call.wait.Done()

    // 从map中删除
    delete(controller.m, accessKey)

    result = call.result
    err = call.err

    controller.mutex.Lock()
    if call.notUse() {
        // 如果没有人使用了,那么就将其放回到临时等待池
        controller.putCall(call)
    }
    controller.mutex.Unlock()

    return
}

func main() {

    testKey:="123"

    var resultA,resultB,resultC interface{}


    go func() {
        resultA,_=Access(testKey, func() (interface{}, error) {
            time.Sleep(time.Second)
            return 1 ,nil
        })
    }()

    go func() {
        resultB,_=Access(testKey, func() (interface{}, error) {
            time.Sleep(time.Second)
            return 2 ,nil
        })
    }()

    go func() {
        resultC,_=Access(testKey, func() (interface{}, error) {
            time.Sleep(time.Second)
            return 3 ,nil
        })
    }()

    time.Sleep(time.Second*2)

    fmt.Println(resultA,resultB,resultC)



}

