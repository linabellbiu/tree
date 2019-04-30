package main

import (
	"fmt"
	"time"
)

type Btree struct {
	key   int
	right *Btree
	left  *Btree
}

type Root struct {
	root  *Btree
	count int
}

func NewBtree(a int) *Btree {
	return &Btree{key: a, right: nil, left: nil}
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

	root := &Root{
		root: nil,
	}

	for _, i := range a {
		root.createBtee(i)
	}

	param := 9999

	// 开始时间
	bT := time.Now()

	//二叉树查找
	fmt.Println("btree查找次数:",root.btreeFind(root.root, param))

	// 从开始到当前所消耗的时间
	eT := time.Since(bT)

	fmt.Println("Run time: ", eT)

	// 开始时间
	bT = time.Now()

	//for循环查找
	fmt.Println("for查找次数:",forFind(param, a))

	// 从开始到当前所消耗的时间
	eT = time.Since(bT)

	fmt.Println("Run time: ", eT)
}

//创建树
func (root *Root) createBtee(a int) {
	btree := NewBtree(a)
	if root.root == nil {
		root.root = btree
	} else {
		createLeaf(root.root, btree)
	}
}

//创建叶子
func createLeaf(oldBtee, newBtree *Btree) {
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

func (root *Root) btreeFind(btree *Btree, param int) int {
	root.count++
	//fmt.Println(root.key)
	if btree.key > param {
		return root.btreeFind(btree.left, param)
	} else if btree.key < param {
		return root.btreeFind(btree.right, param)
	} else if btree.key == param {

	}
	return root.count
}

func forFind(param int, a []int) int {

	//循环的次数
	var count int
	for _, k := range a {
		count++
		if k == param {
			break
		}
	}
	return count
}

func echo(root *Btree) {
	if root.left != nil {
		echo(root.left)
	}
	fmt.Print(root.key)
	fmt.Print(" ")
	if root.right != nil {
		echo(root.right)
	}
}
