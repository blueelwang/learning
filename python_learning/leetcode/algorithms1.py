class Solution(object):
    """
    Given an array of integers, return indices of the two numbers such that they add up to a specific target.
    You may assume that each input would have exactly one solution.
    Example:
    Given nums = [2, 7, 11, 15], target = 9,
    Because nums[0] + nums[1] = 2 + 7 = 9,
    return [0, 1].
    The return format had been changed to zero-based indices.
    """

    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        keymap = {}
        for i in range(len(nums)):
            key = str(target - nums[i])
            if (key in keymap):
                return [keymap[key], i]
            keymap[str(nums[i])] = i

solution = Solution()
solution.twoSum(1, 2)


