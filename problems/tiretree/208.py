# 208. Implement Trie (Prefix Tree)
# Medium
#
# 4046
#
# 64
#
# Add to List
#
# Share
# Implement a trie with insert, search, and startsWith methods.
#
# Example:
#
# Trie trie = new Trie();
#
# trie.insert("apple");
# trie.search("apple");   // returns true
# trie.search("app");     // returns false
# trie.startsWith("app"); // returns true
# trie.insert("app");
# trie.search("app");     // returns true

class TreeNode:
    def __init__(self, val):
        self.val = val
        self.subs = dict()
        self.isWord = False


class Trie:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.root = TreeNode('0')

    def insert(self, word: str) -> None:
        """
        Inserts a word into the trie.
        """
        cur = self.root

        for s in word:
            if s not in cur.subs:
                cur.subs[s] = TreeNode(s)
                cur = cur.subs[s]
            else:
                cur = cur.subs[s]
        cur.isWord = True

    def search(self, word: str) -> bool:
        """
        Returns if the word is in the trie.
        """
        cur = self.root
        for s in word:
            if s in cur.subs:
                cur = cur.subs[s]
            else:
                return False
        return cur.isWord

    def startsWith(self, prefix: str) -> bool:
        """
        Returns if there is any word in the trie that starts with the given prefix.
        """
        cur = self.root
        for s in prefix:
            if s in cur.subs:
                cur = cur.subs[s]
            else:
                return False
        return True

# Your Trie object will be instantiated and called as such:
# obj = Trie()
# obj.insert(word)
# param_2 = obj.search(word)
# param_3 = obj.startsWith(prefix)