package main

func main() {
	twoSum([]int{1, 2, 3, 4}, 5)
	// var results []int = twoSum([]int{1, 2, 3, 4}, 5)
	// for i := range results {
	// 	println(results[i])
	// }
}

func twoSum(nums []int, target int) []int {
	var results = []int{}

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				results = append(results, i, j)
				break
			}
		}

		if len(results) > 0 {
			break
		}
	}

	return results
}
