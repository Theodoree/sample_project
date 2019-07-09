package _00_799

import "fmt"

//哈希表设计
type MyHashMap struct {
    Head *node
}

type node struct {
    Left  *node
    Right *node
    key   int
    value int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
    return MyHashMap{}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
    n := &node{key: key, value: value}
    if this.Head == nil {
        this.Head = n
    } else {
        this.Head.Insert(n)
    }
}

func (n *node) Insert(t *node) {
    current := n

    for {

        if current.key > t.key {
            if current.Left != nil {
                current = current.Left
            } else {
                current.Left = t
                break
            }
        } else if current.key < t.key {
            if current.Right != nil {
                current = current.Right
            } else {
                current.Right = t
                break
            }
        } else if current.key == t.key {
            current.value = t.value
            break
        }
    }
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
    return this.Head.Find(key)

}
func (n *node) Find(key int) int {
    current := n

    for {
        if current == nil {
            return -1
        }
        if current.key > key {
            current = current.Left
        } else if current.key < key {
            current = current.Right
        } else if current.key == key {
            return current.value
        }

    }

}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
    this.Head.Remove(key)

}
func (r *node) Remove(target int) {
    delete(r, target)
}
func delete(node *node, target int) *node {
    if node == nil {
        return nil
    }

    if node.key == target {
        minNode, minNodeParent := min(node.Right)
        minNodeParent.Left = nil
        minNode.Left = node.Left
        minNode.Right = node.Right

        return minNode
    }

    if node.key > target {
        node.Right = delete(node.Right, target)
        fmt.Printf("%#v \n",node)

        return node
    }

    if node.key < target {
        node.Left = delete(node.Left, target)

        return node
    }

    return nil
}

func min(n *node) (*node, *node) {
    current := n
    var parten *node
    for {
        if current.Left != nil {
            parten = current
            current = current.Left
        } else {
            return current, parten
        }
    }
}

func main() {
    /*
       ["MyHashMap","put","put","get","get","put","get", "remove", "get"]
       [[],[1,1],[2,2],[1],[3],[2,1],[2],[2],[2]]
    */

    m := Constructor()
    m.Put(1, 1)
    m.Put(2, 2)
    fmt.Println(m.Get(1))
    fmt.Println(m.Get(3))
    m.Put(2, 1)
    fmt.Println(m.Get(2))
    m.Remove(2)
    fmt.Println(m.Get(2))

}

