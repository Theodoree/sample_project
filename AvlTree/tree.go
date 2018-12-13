package main

import (
	"fmt"
	"math"
)

type AvlBitTree struct {
	data   int
	lchild *AvlBitTree
	rchild *AvlBitTree
	bt     int
}

type IAvlBitTree interface {
	Less(int) bool
	Equal(int) bool
}

func (this *AvlBitTree) Less(data int) bool {
	return data < this.data
}

func (this *AvlBitTree) Equal(data int) bool {
	return data == this.data
}

func GetAvlHeight(tree *AvlBitTree) int {
	if tree == nil {
		return 0
	}

	return int(math.Max(float64(GetAvlHeight(tree.lchild)), float64(GetAvlHeight(tree.rchild)))) + 1
}

func DeleteAvlTree(tree *AvlBitTree, data int) *AvlBitTree {
	tree1 := &tree
	delteavl(tree, tree, tree1, data)
	return *tree1
}

func InsertAvlBitTree(tree *AvlBitTree, data int) *AvlBitTree {
	tree1 := &tree
	insertavl(tree, &tree, data)
	return *tree1
}

func R_Rorate(tree *AvlBitTree, root **AvlBitTree) *AvlBitTree {
	p := tree.lchild
	tree.lchild = p.rchild
	p.rchild = tree
	p.bt = GetAvlHeight(p.lchild) - GetAvlHeight(p.rchild)
	tree.bt = GetAvlHeight(tree.lchild) - GetAvlHeight(tree.rchild)
	if tree == *root {
		*root = p
	}
	tree = p
	return tree
}

func L_Rorate(tree *AvlBitTree, root **AvlBitTree) *AvlBitTree {
	p := tree.rchild
	tree.rchild = p.lchild
	p.lchild = tree
	p.bt = GetAvlHeight(p.lchild) - GetAvlHeight(p.rchild)
	tree.bt = GetAvlHeight(tree.lchild) - GetAvlHeight(tree.rchild)
	if tree == *root {
		*root = p
	}
	tree = p
	return tree
}

func LR_Rorate(tree *AvlBitTree, root **AvlBitTree) *AvlBitTree {
	tree.lchild = L_Rorate(tree.lchild, root)
	return R_Rorate(tree, root)
}

func RL_Rorate(tree *AvlBitTree, root **AvlBitTree) *AvlBitTree {
	tree.rchild = R_Rorate(tree.rchild, root)
	return L_Rorate(tree, root)
}

func MiddleAvlPrint(tree *AvlBitTree) {
	if tree != nil {
		MiddleAvlPrint(tree.lchild)
		fmt.Println(tree.data, tree)
		MiddleAvlPrint(tree.rchild)
	}
}

func blanceavl(tree *AvlBitTree, root **AvlBitTree, data int) {
	//fmt.Println(tree)
	tree.bt = GetAvlHeight(tree.lchild) - GetAvlHeight(tree.rchild)
	if tree.bt >= 2 {
		if data < tree.lchild.data {
			R_Rorate(tree, root)
		} else {
			LR_Rorate(tree, root)
		}
	}

	if tree.bt <= -2 {
		if data >= tree.rchild.data {
			L_Rorate(tree, root)
		} else {
			RL_Rorate(tree, root)
		}
	}
}

func delteavl(tree *AvlBitTree, parent *AvlBitTree, root **AvlBitTree, data int) *AvlBitTree {
	if tree != nil {
		if tree.Equal(data) {
			if *root == tree {
				if tree.lchild == nil {
					*root = tree.rchild
					parent = *root
				} else if tree.rchild == nil {
					*root = tree.lchild
					parent = *root
				} else {
					if tree.lchild.rchild == nil {
						tree.lchild.rchild = tree.rchild
						*root = tree.lchild
						parent = *root
					} else {
						tree1, tree2 := tree.lchild, tree.lchild
						for ; tree1.rchild != nil; {
							tree2 = tree1
							tree1 = tree1.rchild
						}

						if tree1 != tree2 {
							if tree1.lchild != nil {
								tree2.rchild = tree1.lchild
							}

							tree2.rchild = nil
							tree1.lchild = tree.lchild
							tree1.rchild = tree.rchild
							*root = tree1
							parent = *root
						} else {
							tree1.rchild = tree.rchild
							*root = tree1
							parent = *root
						}
					}
				}
			} else if tree.lchild == nil {
				if parent.Less(data) {
					parent.lchild = tree.rchild
				} else {
					parent.rchild = tree.rchild
				}
			} else if tree.rchild == nil {
				if parent.Less(data) {
					parent.lchild = tree.rchild
				} else {
					parent.rchild = tree.rchild
				}
			} else {
				parent.lchild = tree.lchild
				tree.lchild.rchild = tree.rchild
			}
		} else if tree.Less(data) {
			delteavl(tree.lchild, tree, root, data)
		} else {
			delteavl(tree.rchild, tree, root, data)
		}
		blanceavl(tree, root, data)
	}
	return tree
}

func insertavl(tree *AvlBitTree, root **AvlBitTree, data int) *AvlBitTree {
	if tree == nil {
		tree = &AvlBitTree{}
		tree.lchild, tree.rchild = nil, nil
		tree.data = data
		tree.bt = 0
		if *root == nil {
			*root = tree
		}
		return tree
	}

	if tree.Less(data) {
		tree.lchild = InsertAvlBitTree(tree.lchild, data)
		tree.bt = GetAvlHeight(tree.lchild) - GetAvlHeight(tree.rchild)
	} else {
		tree.rchild = InsertAvlBitTree(tree.rchild, data)
		tree.bt = GetAvlHeight(tree.lchild) - GetAvlHeight(tree.rchild)
	}

	blanceavl(tree, root, data)
	return tree
}
