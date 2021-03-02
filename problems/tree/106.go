package main

func main() {

}

func buildTree106(inorder []int, postorder []int) *TreeNode {
    if len(postorder) == 0 {
        return nil
    }
    n := len(postorder)
    node := &TreeNode{Val: postorder[n - 1]}
    index := 0
    for i := index; i < len(inorder); i++ {
        if inorder[i] == node.Val {
            index = i
            break
        }
    }
    node.Left = buildTree(inorder[:index], postorder[:index])
    node.Right = buildTree(inorder[index+1:], postorder[index: n-1])
    return node
}
