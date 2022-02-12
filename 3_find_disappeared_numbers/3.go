package find_disappeared_numbers

// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

// O(n) time complexity, O(n) space complexity
func findDisappearedNumbers2(nums []int) []int {
	unique := make(map[int]bool)
	for i := 1; i <= len(nums); i++ {
		unique[i] = true
	}

	for _, num := range nums {
		if unique[num] {
			delete(unique, num)
		}
	}

	disappeared := make([]int, len(unique))
	i := 0
	for num := range unique {
		disappeared[i] = num
		i++
	}

	return disappeared
}

// O(n) time complexity, O(1) space complexity
func findDisappearedNumbers(nums []int) []int {
	// По условию числа от 1 до n, а индексы от 0 до n-1.
	// Для каждого числа num пометим индекс num-1 с помощью отрицания числа
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		index := Abs(num) - 1

		if nums[index] > 0 {
			nums[index] *= -1
		}
	}

	// Обратное преобразование.
	// Положительный num означает, что элемент равный i+1 не был в исходном массиве
	var disappeared []int
	for i, num := range nums {
		if num > 0 {
			disappeared = append(disappeared, i+1)
		}
	}

	return disappeared
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
