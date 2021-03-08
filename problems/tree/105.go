package main

func main(){

}
type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
  if len(preorder) == 0 {
    return nil
  }
  if len(preorder) == 1 {
    return &TreeNode{
      Val: preorder[0],
    }
  }
  index := 0
  for i, v := range inorder {
    if v == preorder[0] {
      index = i
      break
    }
  }
  return &TreeNode{
    Val: preorder[0],
    Left: buildTree(preorder[1:index+1], inorder[:index]),
    Right: buildTree(preorder[index+1:], inorder[index+1:]),
  }
}
