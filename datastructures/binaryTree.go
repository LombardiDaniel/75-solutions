package datastructures

import "fmt"

type BinaryTree struct {
	Root 	*treeNode
}

type treeNode struct {
	Val		int
	L 		*treeNode
	R 		*treeNode
}

func NewBinaryTree() BinaryTree {
	return BinaryTree{
		Root: nil,
	}
}

func insertOnNode(node *treeNode, v int) *treeNode {
	// fmt.Println(n, v)
	if node == nil {
		newNode := &treeNode{
			Val: v,
			L: nil,
			R: nil,
		}
		return newNode
	}

	if v < node.Val { // go left
		node.L = insertOnNode(node.L, v)
	} else { // go right
		node.R = insertOnNode(node.R, v)
	}

	return node
}

func (tree *BinaryTree) Insert(v int) {
	if tree.Root == nil {
		tree.Root = &treeNode{
			Val: v,
			L: nil,
			R: nil,
		}
		return
	}

	fmt.Println("starting recursive")
	if v < tree.Root.Val { // go left
		tree.Root.L = insertOnNode(tree.Root.L, v)
	} else { // go right
		tree.Root.R = insertOnNode(tree.Root.R, v)
	}
}