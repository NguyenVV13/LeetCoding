package main

func main() {
	twoSumBrute([]int{1, 2, 3, 4}, 5)
	// var results []int = twoSum([]int{1, 2, 3, 4}, 5)
	// for i := range results {
	// 	println(results[i])
	// }
	twoSumMap([]int{1, 2, 3, 4}, 5)
}

func twoSumBrute(nums []int, target int) []int {
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

func twoSumMap(nums []int, target int) []int {
	var results = []int{}
	var differences = make(map[int]int)

	for i, value := range nums {
		if index, ok := differences[value]; ok {
			results = append(results, index, i)
			break
		}

		differences[target-value] = i
	}

	return results
}
