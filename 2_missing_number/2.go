package missing_number

// https://leetcode.com/problems/missing-number/

// O(n) time complexity, O(1) space complexity
func missingNumber(nums []int) int {
	max := len(nums)
	lenPlusMissing := len(nums) + 1
	// Арифм прогрессия. Сумма целых чисел является целым числом
	sumPlusMissing := (0 + max) * lenPlusMissing / 2

	return sumPlusMissing - sum(&nums)
}

func sum(nums *[]int) (sum int) {
	for _, num := range *nums {
		sum += num
	}

	return
}
