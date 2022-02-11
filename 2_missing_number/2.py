from typing import List

# https://leetcode.com/problems/contains-duplicate/


class Solution:
    # O(n) time complexity, O(1) space complexity
    def missingNumber(self, nums: List[int]) -> int:
        max = len(nums)
        len_plus_missing = len(nums) + 1
        # Арифм прогрессия. Сумма целых чисел является целым числом
        sum_plus_missing = (0 + max) * len_plus_missing // 2

        return sum_plus_missing - sum(nums)