
func main() {

}

func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(postorder) == 0 {
        return nil
    }
    l := len(postorder)
    root := &TreeNode{Val: postorder[l-1]}
    index := 0
    for i, v := range inorder {
        if v == root.Val {
            index = i
            break
        }
    }
    root.Left = buildTree(inorder[:index], postorder[:index])
    root.Right = buildTree(inorder[index+1:], postorder[index:l-1])
    return root
}
