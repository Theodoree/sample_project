package fdp

import "sync"

//  Facade Design Pattern

// 加入购物车,并且站内提示

type Handler interface {
    AddSku(userId uint64, Sku uint64)
}

type service struct {
}

func (s *service) AddSku(userId uint64, Sku uint64) {

    GetCartService().Add(userId, Sku)
    createCartNotify(userId, Sku).send()

}


// 提示
type notify interface {
    send()
}

type cartNotify struct {
    UserId uint64 `json:"user_id"`
    SkuId  uint64 `json:"sku_id"`
}

func (c *cartNotify) send() {
    // do something
}

func createCartNotify(userId uint64, Sku uint64) notify {
    return &cartNotify{
        UserId: userId,
        SkuId:  Sku,
    }
}




// 购物车
type Cart interface {
    Add(userId uint64, Sku uint64) error
}

var cartService cart

type cart struct {
    init uint8
    sync.Once
}

func (c *cart) Add(userId uint64, Sku uint64) error {
    // do something
    return nil
}

func GetCartService() Cart {
    if cartService.init == 0 {
        cartService.Do(func() {
            //init
        })
    }
    return &cartService
}
