package main

import (
	"fmt"
	"os"
	"time"
)

type btree struct {
	key   int
	right *btree
	left  *btree
}

type root struct {
	root *btree
}

func NewBtree(a int) *btree {
	return &btree{key: a, right: nil, left: nil}
}

func main() {
	m := make(map[int]int)
	var a []int
	for i := 0; i < 10000000; i++ {
		m[i] = i
	}

	for _, j := range m {
		a = append(a, j)
	}

	root := &root{
		root: nil,
	}

	for _, i := range a {
		root.createBtee(i)
	}

	bT := time.Now() // 开始时间
	find(root.root, 99999)
	eT := time.Since(bT) // 从开始到当前所消耗的时间
	fmt.Println("Run time: ", eT)

	bT = time.Now() // 开始时间

	var o int
	for _, k := range a {
		o++
		if k == 998 {
			eT = time.Since(bT) // 从开始到当前所消耗的时间
			fmt.Println("Run time: ", eT)
			fmt.Println(o)
			os.Exit(1)
		}
	}
}

//创建树
func (root *root) createBtee(a int) {
	btree := NewBtree(a)
	if root.root == nil {
		root.root = btree
	} else {
		createLeaf(root.root, btree)
	}
}

//创建叶子
func createLeaf(oldBtee, newBtree *btree) {
	if oldBtee.key < newBtree.key {
		if oldBtee.right == nil {
			oldBtee.right = newBtree
		} else {
			createLeaf(oldBtee.right, newBtree)
		}
	} else {
		if oldBtee.left == nil {
			oldBtee.left = newBtree
		} else {
			createLeaf(oldBtee.left, newBtree)
		}
	}
}

func find(root *btree, param int) int {
	fmt.Println(root.key)
	if root.key > param {
		return find(root.left, param)
	} else if root.key < param {
		return find(root.right, param)
	} else if root.key == param {
		return root.key
	}
	return 0
}

func echo(root *btree) {
	if root.left != nil {
		echo(root.left)
	}
	fmt.Print(root.key)
	fmt.Print(" ")
	if root.right != nil {
		echo(root.right)
	}
}
