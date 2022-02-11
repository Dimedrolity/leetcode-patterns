package contains_duplicate

//также можно использовать готовый set
//import "github.com/deckarep/golang-set"

//https://leetcode.com/problems/contains-duplicate/

// O(n) time complexity, O(n) space complexity
func containsDuplicate(nums []int) bool {
	type void struct{}

	unique := make(map[int]void)
	for _, num := range nums {
		if _, exists := unique[num]; exists {
			return true
		}

		unique[num] = void{}
	}

	return false
}
