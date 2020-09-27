package rbtree

/*
https://zh.wikipedia.org/wiki/%E7%BA%A2%E9%BB%91%E6%A0%91
*/

/*
   rule-1:节点是红色或黑色。
   rule-2:根是黑色。
   rule-3:所有叶子都是黑色（叶子是NIL节点）。
   rule-4:每个红色节点必须有两个黑色的子节点。（从每个叶子到根的所有路径上不能有两个连续的红色节点。）
   rule-5:从任一节点到其每个叶子的所有简单路径都包含相同数目的黑色节点。
*/

type RBTree struct {
    root *node
    size uint64
}

/*
desc:返回size
*/
func (tree *RBTree) Size() uint64 {
    return tree.size
}

/*
desc:插入k、v
*/
func (tree *RBTree) Insert(key, value interface{}) {}

/*
desc:删除给定key的节点
*/
func (tree *RBTree) Delete(key interface{}) bool {}

/*
desc:返回给定Key的value
*/
func (tree *RBTree) Find(key interface{}) (interface{}, bool) {}

/*
desc:左旋
*/
func (tree *RBTree) rotateLeft(n *node)                       {}

/*
desc:右旋
*/
func (tree *RBTree) rotateRight(n *node)                      {}
