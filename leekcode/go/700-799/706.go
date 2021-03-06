package _00_799

import "fmt"


/*
706. 设计哈希映射

不使用任何内建的哈希表库设计一个哈希映射

具体地说，你的设计应该包含以下的功能

put(key, value)：向哈希映射中插入(键,值)的数值对。如果键对应的值已经存在，更新这个值。
get(key)：返回给定的键所对应的值，如果映射中不包含这个键，返回-1。
remove(key)：如果映射中存在这个键，删除这个数值对。

示例：

MyHashMap hashMap = new MyHashMap();
hashMap.put(1, 1);
hashMap.put(2, 2);
hashMap.get(1);            // 返回 1
hashMap.get(3);            // 返回 -1 (未找到)
hashMap.put(2, 1);         // 更新已有的值
hashMap.get(2);            // 返回 1
hashMap.remove(2);         // 删除键为2的数据
hashMap.get(2);            // 返回 -1 (未找到)

注意：

所有的值都在 [1, 1000000]的范围内。
操作的总数目在[1, 10000]范围内。
不要使用内建的哈希库。
在真实的面试中遇到过这道题？
*/
type MyHashMap struct {
    arr    []int
    IsHave []bool
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
    return MyHashMap{arr: make([]int, 100000), IsHave: make([]bool, 100000)}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
    this.arr[key] = value
    this.IsHave[key] = true
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
    if this.IsHave[key] {
        return this.arr[key]
    }

    return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
    this.IsHave[key] = false
    this.arr[key] = 0

}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

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

