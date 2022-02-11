package __contains_duplicate

// https://leetcode.com/problems/contains-duplicate/

// O(n) time complexity, O(n) space complexity
func containsDuplicate(nums []int) bool {
	unique := make(map[int]bool)
	for _, num := range nums {
		if unique[num] {
			return true
		}

		unique[num] = true
	}

	return false
}
