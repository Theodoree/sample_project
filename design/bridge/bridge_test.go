package bridge

import "testing"

/*
   tip:将抽象和实现解耦，让它们可以独立变化
   tip:
    使用场景:
        一个功能可能会有多种实现类,比如说支付,钱包支付 移动支付 银行卡支付。
        数据库 mysql mongo tidb .....
*/

func TestNewService(t *testing.T) {
    s, err := NewService("mongo:12312312")
    if err != nil {
        t.Fatal(err)
    }
    s.Open()

    s, err = NewService("mysql:12312312")
    if err != nil {
        t.Fatal(err)
    }
    s.Open()

    s, err = NewService("oracle:12312312")
    if err != nil {
        t.Fatal(err)
    }
    s.Open()

    s, err = NewService("123:12312312")
    if err != nil {
        t.Fatal(err)
    }
    s.Open()
}
