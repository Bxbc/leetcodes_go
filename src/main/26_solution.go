package main

import "fmt"

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/?envType=study-plan-v2&envId=top-interview-150
func removeDuplicates(nums []int) int {
	i, j := 0, 1
	for {
		if i >= len(nums) || j >= len(nums) {
			break
		}
		if nums[i] == nums[j] {
			i++
			j++
		} else {
			var temp = nums[j]
			nums[j] = nums[i]
			nums[i] = temp
			i++
		}
	}
	return i
}

func main() {
	nums := []int{1, 1, 2}
	duplicates := removeDuplicates(nums)

	fmt.Println(duplicates, nums)
}
