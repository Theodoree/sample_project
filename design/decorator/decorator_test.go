package decorator

import (
    "testing"
)

/*
tip:如果设计者不需要用户关 注是否使用缓存功能，要隐藏实现细节，也就是说用户只能看到和使用代理类，那么就使 用proxy模式
    ;反之，如果设计者需要用户自己决定是否使用缓存的功能，需要用户自己新 建原始对象并动态添加缓存功能，那么就使用decorator模式。

*/

func TestNewThreadPool(t *testing.T) {
    pool := NewThreadPool(10)

    for i:=0;i<10000;i++{
        pool.SendNewFunc(func() {
            println(1)
        })
    }

}
