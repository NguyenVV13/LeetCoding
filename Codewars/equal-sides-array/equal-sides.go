package main

func main() {
	// var arr1 = []int{1, 2, 3, 4, 3, 2, 1}
	var arr2 = []int{1, 100, 50, -51, 1, 1}
	println(FindEvenIndex(arr2))
}

// Time complexity: O(n^2)
// Slight improvement would be to start in the middle and
// move in the direction of the smaller absolute sum
func FindEvenIndex(arr []int) int {
	for i := range arr {
		if calcSum(arr[:i]) == calcSum(arr[i+1:]) {
			return i
		}
	}
	return -1
}

func calcSum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}
