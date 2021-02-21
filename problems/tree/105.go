
func main(){

}
type TreeNode struc {
  Val int
  Lett *TreeNode
  Right *TreeNode
}

func buildTree(pre []int, ord []int) *TreeNode {
  if len(pre) == 0 {
    returen None
  }
  root = &TreeNode{Val: pre[0]}
  index := 0
  for i, v := range ord {
    if v == pre[0] {
      index = i
      break
    }
  }
  root.Left = buildTree(pre[1:index+1], ord[:index])
  root.Right = buildTree(pre[index+1:], ord[index+1:])
  return root
}
