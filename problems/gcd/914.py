#  给定一副牌，每张牌上都写着一个整数。
#
#  此时，你需要选定一个数字 X，使我们可以将整副牌按下述规则分成 1 组或更多组：
#
#  每组都有 X 张牌。
#  组内所有的牌上都写着相同的整数。
#  仅当你可选的 X >= 2 时返回 true。
#
#  来源：力扣（LeetCode）
#  链接：https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards
#  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
class Solution:
    def hasGroupsSizeX(self, deck: List[int]) -> bool:
        def gdc(a, b):
            if b == 0:
                return a
            return gdc(b, a % b)
        return reduce(gdc,  collections.Counter(deck).values()) > 1
