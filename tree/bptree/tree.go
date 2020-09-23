package bptree

import (
    "container/list"
    "fmt"
    "unsafe"
)

const (
    DEFAULT_ORDER = 1 << 10 // 1024 2048 4096
    // 每次节点都有order个节点
    order = DEFAULT_ORDER
)


/*tip: tree
                                                    root
                                        child[node_a,node_b,node_c]
                            /                          |                        \
                        nodea                       node_b                      node_c
    child[value_a,value_b....  node_b]    child[value_c value_d ...... node_c]   child[value_e value_f ....... 0]
*/




/*
tip:
    规则A:最后一个child为node节点,指向下一个节点
*/

type BpTree struct {
    root *node
    pool *resourcePool // 8byte
}

/*
tip: <<<<<<<<<<<<<<-打印相关函数->>>>>>>>>>>>>
*/

func (tree *BpTree) pathToRoot(child *node) uint64 {
    c := child
    var length uint64
    for c != tree.root {
        c = c.Parent
        length++
    }
    return length

}
func (tree *BpTree) height() uint64 {
    var cnt uint64
    root := tree.root
    for root != nil {
        root = (*node)(root.child[0])
        cnt++
    }
    return cnt
}
/*
desc: 打印所有叶子节点
*/
func (tree *BpTree) PrintLeaves() {
    if tree.root == nil {
        return
    }

    cur := tree.root

    // 递归找到叶子节点
    for !cur.IsLeaf {
        cur = (*node)(cur.child[0])
    }

    for {
        for i := 0; i < cur.numKeys; i++ {
            fmt.Println(cur.childKeys[i],(*value)(cur.child[i]).val)
        }


        /*
        case:如果child[order-1]等于nil的话,那么说明已经遍历完毕,否则继续获取下一个叶子节点
        */
        if cur.child[order-1] != nil {
            cur = (*node)(cur.child[order-1])
        } else {
            break
        }
    }
    fmt.Println()
}
/*
desc: 打印整个tree
*/
func (tree *BpTree) PrintTree() {

    root := tree.root
    if root == nil {
        return
    }

    var (
        depth uint64
        l     = list.New()
    )

    // 先进先出
    l.PushBack(root)
    for l.Len() != 0 {
        n := l.Front().Value.(*node)
        if n.Parent != nil && unsafe.Pointer(n) == n.Parent.child[0] {
            curDepth := tree.pathToRoot(n)
            if curDepth != depth {
                depth = curDepth // 更新深度
                println()        // 换行
            }
        }

        for i := 0; i < n.numKeys; i++ {
            print(n.childKeys[i])
        }
        if !n.IsLeaf {
            for i := 0; i < n.numKeys; i++ {
                l.PushBack((*node)(n.child[i])) // 入队
            }
        }
        print("|  ")

    }
    println()
}

/*
desc: 打印单个value节点
*/
func (tree *BpTree) findAndPrint(key uint64) {
    val := tree.find(key, nil)
    if val != nil {
        fmt.Printf("%p %d %v \n", val, key, val.val)
    } else {
        println("not found")
    }
}

/*
desc: 范围遍历
*/
func (tree *BpTree) findAndPrintRange(start, end uint64) {
    key, arr := tree.findRange(start, end)
    if len(arr) > 0 {
        for i := 0; i < len(arr); i++ {
            fmt.Printf("Key: %d   Location: %p  Value: %d \n", key[i], arr[i], arr[i].val)
        }
    }
}

/*
tip: <<<<<<<<<<<<<<-查找相关函数->>>>>>>>>>>>>
*/

// 返回value
func (tree *BpTree) find(key uint64, leafOut **node) *value {
    if tree.root == nil {
        return nil
    }

    var (
        i    int
        leaf = tree.findLeaf(key)
    )

    for i = 0; i < leaf.numKeys; i++ {
        if leaf.childKeys[i] == key {
            break
        }
    }

    if leafOut != nil {
        *leafOut = leaf
    }

    if i < leaf.numKeys {
        return (*value)(leaf.child[i])
    }

    return nil
}

// 返回Start-end之间的节点
func (tree *BpTree) findRange(start, end uint64) ([]uint64, []*value) {
    leaf := tree.findLeaf(start)
    if leaf == nil {
        return nil, nil
    }
    var i int
    for i = 0; i < leaf.numKeys && leaf.childKeys[i] < start; i++ {
    }

    if i == leaf.numKeys {
        return nil, nil
    }

    var arr []*value
    var key []uint64

    for leaf != nil {
        for ; i < leaf.numKeys && leaf.childKeys[i] <= end; i++ {
            key = append(key, leaf.childKeys[i])
            arr = append(arr, (*value)(leaf.child[i]))
        }
        leaf = (*node)(leaf.child[order-1]) // 获取下一个叶子节点
        i = 0
    }

    return key, arr
}

// 返回叶子节点
func (tree *BpTree) findLeaf(key uint64) *node {
    if tree.root == nil {
        return nil
    }

    n := tree.root

    for !n.IsLeaf {
        i := 0

        for i < n.numKeys { // 寻找位置
            if key >= n.childKeys[i] {
                i++
            } else {
                break
            }
        }
        n = (*node)(n.child[i])
    }

    return n
}

// 返回node在parent subChild中的下标
func (tree *BpTree) GetLeftIndex(parent, left *node) int {
    var leftIndex = 0
    for leftIndex <= parent.numKeys && parent.child[leftIndex] != unsafe.Pointer(left) {
        leftIndex++
    }
    return leftIndex
}

// 返回邻居节点下标 如果没有邻居那么是返回-1
func (tree *BpTree) getNeighborIndex(n *node) (int, error) {

    for i := 0; i < n.Parent.numKeys; i++ {
        if n.Parent.child[i] == unsafe.Pointer(n) {
            return i - 1, nil
        }
    }

    return 0, notFoundNeighborError
}

/*
tip: <<<<<<<<<<<<<<-插入相关函数->>>>>>>>>>>>>
*/

// 将k、v 插入叶子节点
func (tree *BpTree) insertIntoLeaf(leaf *node, key uint64, val *value) *node {
    var (
        insertionPoint int
    )

    if !leaf.IsLeaf {
        return nil
    }

    // 寻找插入位置
    for insertionPoint < leaf.numKeys && leaf.childKeys[insertionPoint] < key {
        insertionPoint++
    }

    for i := leaf.numKeys; i > insertionPoint; i-- {
        // 插入位置之后的所有元素向前移动一格
        leaf.childKeys[i] = leaf.childKeys[i-1]
        leaf.child[i] = leaf.child[i-1]
    }

    leaf.childKeys[insertionPoint] = key
    leaf.child[insertionPoint] = unsafe.Pointer(val)
    return leaf
}

// 叶子节点分裂
func (tree *BpTree) insertLeafAfterSplitting(leaf *node, key uint64, record *value) {

    if !leaf.IsLeaf {
        return
    }

    newLeaf := tree.pool.GetLeaf()

    insertionIndex := 0
    // 找到插入位置
    for insertionIndex < order-1 && leaf.childKeys[insertionIndex] < key {
        insertionIndex++
    }

    tmpKey := tree.pool.GetTempSlice()
    tmpChild := tree.pool.GetTempSlice()

    var i, j int
    // 针对高缓优化
    for i < leaf.numKeys {
        // skip 插入位置
        if j == insertionIndex {
            j++
        }
        tmpChild[j] = leaf.child[i]
        tmpKey[j] = leaf.childKeys[i]
    }

    tmpKey[insertionIndex] = key
    tmpChild[insertionIndex] = unsafe.Pointer(record)

    leaf.numKeys = 0
    split := cut(order - 1)

    // 页分裂,重置节点
    for i := 0; i < split; i++ {
        leaf.childKeys[i] = tmpKey[i].(uint64)
        leaf.child[i] = tmpChild[i].(unsafe.Pointer)
        leaf.numKeys++
    }

    i = split
    j = 0
    // j 是newNode的索引
    for i < order {
        newLeaf.childKeys[j] = tmpKey[i].(uint64)
        newLeaf.child[j] = tmpChild[i].(unsafe.Pointer)
        newLeaf.numKeys++
        i++
        j++
    }

    tree.pool.PutTempSlice(tmpKey)
    tree.pool.PutTempSlice(tmpChild)

    // before: oldLeaf      -      oldRightLeaf
    // after: oldLeaf - newLeaf - oldRightLeaf
    newLeaf.child[order-1] = leaf.child[order-1]
    leaf.child[order-1] = unsafe.Pointer(newLeaf)

    // 将分裂出去的child节点设置为空
    for i := leaf.numKeys; i < order-1; i++ {
        leaf.child[i] = nil
    }

    newLeaf.Parent = leaf.Parent

    newKey := newLeaf.childKeys[0]

    tree.insertIntoParent(leaf, newKey, newLeaf)
}

// node节点分裂
func (tree *BpTree) insertNodeAfterSplitting(parent, right *node, leftIndex int, key uint64) {

    tmpChild := tree.pool.GetTempSlice()
    tmpKey := tree.pool.GetTempSlice()

    var i, j int

    for i < parent.numKeys+1 {
        // 插入位置,跳过
        if j == leftIndex+1 {
            j++
        }
        tmpChild[j] = parent.child[i]
        tmpKey[j] = parent.child[i]
    }

    tmpChild[leftIndex+1] = right
    tmpKey[leftIndex] = key

    //开始分裂
    split := cut(order)
    newParentNode := tree.pool.GetNode()
    parent.numKeys = 0
    for i := 0; i < split-1; i++ {
        parent.childKeys[i] = tmpKey[i].(uint64)
        parent.child[i] = tmpChild[i].(unsafe.Pointer)
        parent.numKeys++
    }

    parent.child[split-1] = tmpChild[split-1].(unsafe.Pointer)
    kPrime := tmpKey[split-2].(uint64)

    // j 是newnode的下标
    j = 0
    i = split
    for i < order {
        newParentNode.child[j] = tmpChild[i].(unsafe.Pointer)
        newParentNode.childKeys[j] = tmpKey[i].(uint64)
        newParentNode.numKeys++
        i++
        j++
    }

    // 最后一个节点是链接的节点 这里的i等于order
    newParentNode.child[j] = tmpChild[i].(unsafe.Pointer)

    tree.pool.PutTempSlice(tmpChild)
    tree.pool.PutTempSlice(tmpKey)

    newParentNode.Parent = parent.Parent
    // 遍历子节点
    for i := 0; i < newParentNode.numKeys; i++ {
        (*node)(newParentNode.child[i]).Parent = newParentNode
    }

    // 将新parent节点与祖先节点相关联
    tree.insertIntoParent(parent, kPrime, newParentNode)
}

// 插入到父节点
func (tree *BpTree) insertIntoParent(left *node, key uint64, right *node) {
    parent := left.Parent
    if parent == nil { // 如果为空,俺么插入新root节点
        tree.insertIntoNewRoot(left, right, key)
        return
    }

    leftIndex := tree.GetLeftIndex(parent, left)
    // 如果有位置,那么直接插入
    if parent.numKeys < order-1 {
        tree.insertIntoNode(parent, right, leftIndex, key)
        return
    }

    // 否则就必须页分裂后插入了
    tree.insertNodeAfterSplitting(parent, right, leftIndex, key)
}
func (tree *BpTree) insertIntoNode(parent, right *node, leftIndex int, key uint64, ) {
    for i := parent.numKeys; i > leftIndex; i-- {
        // 整体向前移动一步
        parent.child[i+1] = parent.child[i]
        parent.childKeys[i] = parent.childKeys[i-1]

    }

    parent.child[leftIndex+1] = unsafe.Pointer(right)
    parent.childKeys[leftIndex] = key
    parent.numKeys++

}
func (tree *BpTree) insertIntoNewRoot(left, right *node, key uint64) {
    tree.root = tree.pool.GetNode()
    tree.root.childKeys[0] = key
    tree.root.child[0] = unsafe.Pointer(left)
    tree.root.child[1] = unsafe.Pointer(right)
    tree.root.numKeys++
    left.Parent = tree.root
    right.Parent = tree.root
}

// 创建root节点
func (tree *BpTree) startNewTree(key uint64, value *value) {
    tree.root = tree.pool.GetLeaf()
    tree.root.childKeys[0] = key
    tree.root.child[0] = unsafe.Pointer(value)
    tree.root.numKeys++
}

/*
tip: <<<<<<<<<<<<<<-删除相关函数->>>>>>>>>>>>>
*/

// 从n的subChild中删除deleteNode
func (tree *BpTree) removeEntryFromNode(n *node, entry unsafe.Pointer, key uint64) {

    i := 0
    // 找到node的节点
    for n.childKeys[i] != key {
        i++
    }

    //将要删除的位置往后的元素都向前移动一格
    i++
    for ; i < n.numKeys; i++ {
        n.childKeys[i-1] = n.childKeys[i]
    }

    // 确定需要扫描的数量
    numPointer := n.numKeys
    if !n.IsLeaf {
        numPointer++
    }

    i = 0
    // key是可以重复的
    for n.child[i] != entry {
        i++
    }

    i++
    for ; i < numPointer; i++ {
        n.child[i-1] = n.child[i]
    }

    // 已被覆盖
    n.numKeys--

    if n.IsLeaf {
        for i := n.numKeys; i < order-1; i++ {
            n.child[i] = nil
        }
    } else {
        for i := n.numKeys + 1; i < order; i++ {
            n.child[i] = nil
        }
    }

}

// 删除
func (tree *BpTree) deleteEntry(n *node, entry unsafe.Pointer, key uint64) error {

    tree.removeEntryFromNode(n, entry, key)

    // 如果n是root节点
    if n == tree.root {
        tree.adjustRoot()
        return nil
    }

    var minKeys int
    // node节点多一个
    minKeys = cut(order) - 1
    if n.IsLeaf { // 叶子节点最后一个是链接的node节点
        minKeys = cut(order - 1)
    }

    // 如果n节点数量大于minKeys,那么不需要调整
    if n.numKeys >= minKeys {
        return nil
    }

    // 否则需要调整
    neighborIndex, err := tree.getNeighborIndex(n)
    if err == nil {
        return err
    }

    kPrimeIndex := neighborIndex
    // 如果parent节点是在数组[0]的位置
    if neighborIndex == -1 {
        kPrimeIndex = 0
    }

    kPrime := n.Parent.childKeys[kPrimeIndex]

    var neighbor *node
    if neighborIndex == -1 { // 如果为-1,那么邻居在右边
        neighbor = (*node)(n.Parent.child[1])
    } else {
        neighbor = (*node)(n.Parent.child[neighborIndex])
    }

    /*
       case: 索引节点,cap = order - 1
       case: 叶子节点,那么cap = order
    */
    var capacity int
    capacity = order - 1
    if n.IsLeaf {
        capacity = order
    }

    /*
       case: 如果小于单个节点的容量,那么就合并
    */
    if neighbor.numKeys+n.numKeys < capacity {
        tree.coalesceNode(n, neighbor, neighborIndex, kPrime)
        return nil
    }
    /*
       case: 否则,就调整节点
    */
    tree.redistributeNodes(n, neighbor, neighborIndex, kPrimeIndex, kPrime)
    return nil
}

/*
tip: <<<<<<<<<<<<<<-调整相关函数->>>>>>>>>>>>>
*/
// 调整root
func (tree *BpTree) adjustRoot() *node {

    if tree.root == nil {
        return nil
    }

    // 不为空的root,直接返回
    if tree.root.numKeys > 0 {
        return tree.root
    }

    var newRoot *node

    // 如果不是叶子节点(子节点是value),且子节点不为空,那么提拔第一个为新root
    if !tree.root.IsLeaf {
        newRoot = (*node)(tree.root.child[0])
        tree.root = newRoot
    } else { // 如果是叶子节点,那么说明整个树都是空的
        tree.root = nil
    }

    return newRoot

}

// 合并节点
func (tree *BpTree) coalesceNode(n, neighbor *node, neighborIndex int, kPrime uint64) {

    /*
       case:如果n节点位于数组的[0]的位置,那么neighborIndex等于-1。
            因为bpt的规则是,左边必然比右边小。然后n处于最0的位置,那么这里就需要交换一下指针
            最后的结果就是将neighbor合并到n上。
    */
    if neighborIndex == -1 {
        n, neighbor = neighbor, n
    }

    neighborInsertionIndex := neighbor.numKeys

    if !n.IsLeaf { // 不是叶子节点 直接加在后面
        neighbor.childKeys[neighborInsertionIndex] = kPrime
        neighbor.numKeys++

        var (
            nEnd = n.numKeys
            j    = 0
            i    = neighborInsertionIndex + 1
        )

        // 将n合并到neighbor
        for j < nEnd {
            neighbor.childKeys[i] = n.childKeys[j]
            neighbor.child[i] = n.child[j]
            neighbor.numKeys++
            n.numKeys--

            i++
            j++
        }

        // 指针的数量永远比key多一个
        neighbor.child[i] = n.child[j]

        for i := 0; i < neighbor.numKeys; i++ {
            (*node)(neighbor.child[i]).Parent = neighbor
        }
    } else { // 是叶子节点
        var (
            j = 0
            i = neighborInsertionIndex
        )
        for j < n.numKeys {

            neighbor.childKeys[i] = n.childKeys[j]
            neighbor.child[i] = n.child[j]
            neighbor.numKeys++
            n.numKeys--

            // 下标移动
            i++
            j++
        }

        // 连接底层value的node
        neighbor.child[order-1] = n.child[order-1]
    }

    tree.deleteEntry(n.Parent, unsafe.Pointer(n), kPrime)

}

// 重新分配child节点
func (tree *BpTree) redistributeNodes(n, neighbor *node, neighborIndex int, kPrimeIndex int, kPrime uint64) {

    if neighborIndex != -1 {
        /*
           邻居在n左边
           那么取出邻居最后一对键值对
           设置到n的左边
        */

        if !n.IsLeaf { //不是叶子节点,那么需要保存最后一个key
            n.child[n.numKeys+1] = n.child[n.numKeys]
        }

        // 整体向后移动一格,那么0的位置现在是空的
        for i := n.numKeys; i > 0; i-- {
            n.child[i] = n.child[i-1]
            n.childKeys[i] = n.childKeys[i-1]
        }

        if !n.IsLeaf {
            // 如果是node节点,那么数量比叶子节点多一个键值对
            n.child[0] = neighbor.child[neighbor.numKeys]
            (*node)(n.child[0]).Parent = n
            n.childKeys[0] = kPrime

            neighbor.child[neighbor.numKeys] = nil

            // 重新设置邻居节点在父节点的最大值(因为最后一个节点移动到b节点了)
            n.Parent.childKeys[kPrimeIndex] = neighbor.childKeys[neighbor.numKeys-1]
        } else {
            // 叶子节点比node少一个键值对
            n.child[0] = neighbor.child[neighbor.numKeys-1]         //获取最后一个叶子节点
            n.childKeys[0] = neighbor.childKeys[neighbor.numKeys-1] //获取最后一个叶子节点
            (*node)(n.child[0]).Parent = n                          // 重新设置parent节点

            n.Parent.childKeys[kPrimeIndex] = n.childKeys[0]
            neighbor.child[neighbor.numKeys-1] = nil
        }
    } else {
        /*
           邻居在右边,从右边的邻居取一对键指针,增加到n上
        */

        if !n.IsLeaf {
            n.child[n.numKeys] = neighbor.child[0]
            n.childKeys[n.numKeys] = neighbor.childKeys[0]
            n.Parent.childKeys[kPrimeIndex] = neighbor.childKeys[1]
        } else {
            n.childKeys[n.numKeys] = kPrime
            n.child[n.numKeys+1] = neighbor.child[0]
            (*node)(n.child[n.numKeys+1]).Parent = n
            n.Parent.childKeys[kPrimeIndex] = neighbor.childKeys[0]
        }

        var i int
        // 向前移动一个一格
        for i = 0; i < neighbor.numKeys-1; i++ {
            neighbor.childKeys[i] = neighbor.childKeys[i+1]
            neighbor.child[i] = neighbor.child[i+1]
        }

        if !n.IsLeaf {
            neighbor.child[i] = neighbor.child[i+1]
        }

    }

    n.numKeys++
    neighbor.numKeys--

}

/*
tip: <<<<<<<<<<<<<<-对外暴露函数->>>>>>>>>>>>>
*/

func (tree *BpTree) Delete(key uint64) {

    var leaf *node

    value := tree.find(key, &leaf)
    /*
       case: 如果找到,那么就直接删除
    */
    if value != nil && leaf != nil {
        tree.deleteEntry(leaf, unsafe.Pointer(value), key)

        // 内存池优化
        value.reset()
        tree.pool.valuePool.Put(value)
    }

    return
}

func (tree *BpTree) Insert(key uint64, value interface{}) {

    val := tree.find(key, nil)

    if val != nil { // 如果找到直接更新
        val.val = value
        return
    }

    val = tree.pool.GetValue()
    val.val = value

    /*
       case: root不存在,创建
    */

    if tree.root == nil {
        tree.startNewTree(key, val)
        return
    }

    /*
       case: 根据key寻找对应的叶子节点
    */
    leaf := tree.findLeaf(key)

    /*
       case: 如果叶子节点有足够位置,直接插入
    */
    if leaf.numKeys < order-1 {
        tree.insertIntoLeaf(leaf, key, val)
        return
    }

    /*
       case: 空间不足,就得页分裂了.
    */
    // 页分裂
    tree.insertLeafAfterSplitting(leaf, key, val)
}

/*
desc:实际上将指针置为空,那么gc就可以正常运行了
*/
func (tree *BpTree) DestroyTree() {
    tree.root = nil
    tree.pool = nil
}

func NewBpTree() *BpTree {
    bpTree := &BpTree{
        root: nil,
        pool: defaultResourcePool,
    }
    return bpTree
}
