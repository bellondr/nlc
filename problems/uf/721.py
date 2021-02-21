import collections

class UF:
    def __init__(self):
        self.parent = {}

    def find(self, x):
        self.parent.setdefault(x, x)
        while x != self.parent[x]:
            x = self.parent[x]
        return x
    def union(self, p, q):
        self.parent[self.find(p)] = self.find(q)


class Solution:
    def accountsMerge(self, accounts: [[str]]) -> [[str]]:
        uf = UF()
        email_to_name = {}
        res = collections.defaultdict(list)
        for account in accounts:
            for i in range(1, len(account)):
                email_to_name[account[i]] = account[0]
                if i < len(account) - 1:uf.union(account[i], account[i + 1])
        print(email_to_name)
        for email in email_to_name:
            res[uf.find(email)].append(email)

        print(res)

        return [[email_to_name[value[0]]] + sorted(value) for value in res.values()]


if __name__ == '__main__':
    s = Solution()
    case = [["John","johnsmith@mail.com","john_newyork@mail.com"],["John","johnsmith@mail.com","john00@mail.com"],["Mary","mary@mail.com"],["John","johnnybravo@mail.com"]]
    print(s.accountsMerge(case))
    
