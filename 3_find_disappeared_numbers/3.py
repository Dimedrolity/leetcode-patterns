from typing import List

# https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

class Solution:
    # O(n) time complexity, O(n) space complexity
    def findDisappearedNumbers_1(self, nums: List[int]) -> List[int]:
        unique_numbers = set(nums)
        missing_numbers = []

        for num in range(1, len(nums) + 1):
            if num not in unique_numbers:
                missing_numbers.append(num)

        return missing_numbers

    # O(n) time complexity, O(1) space complexity
    # подсмотрено на ytube, разобрано
    def findDisappearedNumbers(self, nums: List[int]) -> List[int]:
        # По условию числа от 1 до n, а индексы от 0 до n-1.
        # Для каждого числа num пометим индекс num-1 с помощью отрицания числа
        for num in nums:
            i = abs(num) - 1
            if nums[i] > 0:
                nums[i] *= -1

        # Обратное преобразование.
        # Положительный num означает, что элемент равный i+1 не был в исходном массиве
        missing = []
        for i, num in enumerate(nums):
            if num > 0:
                missing.append(i + 1)

        return missing
