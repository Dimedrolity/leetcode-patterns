from typing import List

# https://leetcode.com/problems/contains-duplicate/


class Solution:
    # O(n) time complexity, O(n) space complexity
    def containsDuplicate(self, nums: List[int]) -> bool:
        return len(nums) != len(set(nums))
