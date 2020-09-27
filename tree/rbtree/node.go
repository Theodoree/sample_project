package rbtree

const (
    black = iota
    red
)

type node struct {
    child [2]*node
    key   interface{}
    value interface{}
    color uint8
}

func newNode() *node {
    return &node{}
}

func (n *node) reset() {
    for idx := range n.child {
        n.child[idx] = nil
    }
    n.key = nil
    n.value = nil
    n.color = black
}


