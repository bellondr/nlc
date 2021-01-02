class Rotate():
    def __init__(self):
        pass
    def reverse(self, nums: [], start: int, end: int):
        while start < end:
            nums[start], nums[end] = nums[end], nums[start]
            start += 1
            end -= 1

    def rotate(self, nums: [], k: int) -> None:
        l = len(nums)
        if l == 0:
            return
        k = k % l
        if k == 0:
            return
        self.reverse(nums,0, l-1)
        self.reverse(nums, 0, k-1)
        self.reverse(nums, k, l-1)
        print(nums)


if __name__ == '__main__':
    r = Rotate()
    r.rotate([1,2,3,4,5,6], 4)
    print('===============================')
    r.rotate([1,2,3,4,5,6,7],3)