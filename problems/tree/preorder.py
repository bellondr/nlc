
from problems.tree.tree import TreeNode

class PreOrder():
    def __init__(self):
        pass

    def pre_order(self, node):
        if node is None:
            return
        print(node.val)
        self.pre_order(node.left)
        self.pre_order(node.right)


if __name__ == '__main__':
    p = PreOrder()
    t = TreeNode(1, TreeNode(2, TreeNode(4), TreeNode(5)), TreeNode(3, TreeNode(6), TreeNode(7)))
    p.pre_order(t)