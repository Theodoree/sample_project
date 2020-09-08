package responsibility

import (
    "fmt"
    "testing"
)

/*
tip:
    应用场景:
        过滤敏感词,
*/



func TestNewFilterChain(t *testing.T) {
    chain:=NewFilterChain()


    chain.AddSubFilter(NewAdWordFilter())
    chain.AddSubFilter(NewsexWordFilter())

    postText:=`大扎好,我是渣渣辉,玩游戏无所谓,是兄弟就来砍我`

    fmt.Printf("过滤后:%s \n",chain.DoFilter(postText))



}

