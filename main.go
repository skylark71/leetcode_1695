package main

import "fmt"

func main() {
	array := []int{4, 2, 4, 5, 6}
	fmt.Println(maximumUniqueSubarray(array))
}

func maximumUniqueSubarray(nums []int) int {
	seen := make(map[int]bool)
	left := 0
	sum, maxSum := 0, 0

	for right := 0; right < len(nums); right++ {
		for seen[nums[right]] {
			seen[nums[left]] = false
			sum -= nums[left]
			left++
		}
		seen[nums[right]] = true
		sum += nums[right]

		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}
