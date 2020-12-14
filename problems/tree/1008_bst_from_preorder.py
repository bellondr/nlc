
class Solution():
    def __init__(self):
        pass

    def bstFromPreorder(self, preorder: List[int]) -> TreeNode:
        def dfs(start, end):
            if start > end:
                return None
            if start == end:
                return TreeNode(preorder[start])
            root = TreeNode(preorder[start])
            mid = -1
            for i in range(start + 1, end + 1):
                if preorder[i] > preorder[start]:
                    mid = i
                    break
            if mid == -1:
                root.left = dfs(start+1, end)
            else:
                root.left = dfs(start+1, mid-1)
                root.right = dfs(mid, end)
            return root
        return dfs(0, len(preorder) - 1)

    def bstFromPreorder2(self, preorder: List[int]) -> TreeNode:
        if len(preorder) == 0:
            return None
        temp = root = TreeNode(preorder[0])
        start, size = 1, len(preorder)
        while start < size:
            if preorder[start] < temp.val:
                if temp.left:
                    temp = temp.left
                else:
                    temp.left = TreeNode(preorder[start])
                    start += 1
                    temp = root
            else:
                if temp.right:
                    temp = temp.right
                else:
                    temp.right = TreeNode(preorder[start])
                    start += 1
                    temp = root
        return root
