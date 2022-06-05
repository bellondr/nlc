package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	testCodec()
}
func testCodec() {
	c := Constructor()
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	ser := c.serialize(root)
	fmt.Println(ser)
	tre := c.deserialize(ser)
	printTree(tre)
}
func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	printTree(root.Left)
	printTree(root.Right)
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Codec struct {
	preOrder []string
}

func Constructor() Codec {
	return Codec{
		preOrder: []string{},
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.toPreArray(root)
	s := strings.Join(this.preOrder, ",")
	return s
}


func (this *Codec)toPreArray(root *TreeNode) {
	if root == nil {
		this.preOrder = append(this.preOrder, "#")
		return
	}
	this.preOrder = append(this.preOrder, fmt.Sprintf("%d", root.Val))
	this.toPreArray(root.Left)
	this.toPreArray(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	preOrder := strings.Split(data, ",")
	this.preOrder = preOrder
	return this.buildTreePre()
}

func (this *Codec)buildTreePre() *TreeNode {
	if len(this.preOrder) == 0 {
		return nil
	}
	if this.preOrder[0] == "#" {
		this.preOrder = this.preOrder[1:]
		return nil
	}
	valStr := this.preOrder[0]
	this.preOrder = this.preOrder[1:]
	val, _:= strconv.Atoi(valStr)
	root := &TreeNode{Val: val}
	root.Left = this.buildTreePre()
	root.Right = this.buildTreePre()
	return root
}

func buildTree(preOrder []string) *TreeNode {
	val, _:= strconv.Atoi(preOrder[0])
	root := &TreeNode{Val: val}
	queue := []*TreeNode{
		root,
	}
	for i := 1; i < len(preOrder); {
		n := len(queue)
		for j := 0; j < n; j++ {
			node := queue[j]
			if preOrder[i] != "#" {
				val, _:= strconv.Atoi(preOrder[i])
				node.Left = &TreeNode{Val: val}
				queue = append(queue, node.Left)
			}
			i++
			if preOrder[i] != "#" {
				val, _:= strconv.Atoi(preOrder[i])
				node.Right = &TreeNode{Val: val}
				queue = append(queue, node.Right)
			}
			i++
		}
		queue = queue[n:]
	}
	return root
}
