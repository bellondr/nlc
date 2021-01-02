class TwoSum:
    def __init__(self):
        pass

    def twoSum(self, nums: List[int], target: int) -> List[int]:
        res = {}
        for i in range(len(nums)):
            key = target - nums[i]
            if key in res:
                return [res[key], i]
            res[nums[i]] = i
