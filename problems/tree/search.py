#二叉树

### 标记法
# 使用颜色标记节点的状态，新节点为白色，已访问的节点为灰色。
# 如果遇到的节点为白色，则将其标记为灰色，然后将其右子节点、自身、左子节点依次入栈。
# 如果遇到的节点为灰色，则将节点的值输出。

class TreeNode():
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class SearchBinTree():
    def __init__(self):
        pass

    def search_tree(self, root):
        res = []
        q = [(1, root)]
        while q:
            c, node = q.pop()
            if node is None: continue
            if c == 1:
                q.append((1, node.right))
                q.append((1, node.left))
                q.append((0, node))
            else:
                res.append(node.val)
        return res


if __name__ == '__main__':
    root = TreeNode(1, TreeNode(2, TreeNode(3), TreeNode(4)), TreeNode(9))
    s = SearchBinTree()
    a =  s.search_tree(root)
    print(a)
