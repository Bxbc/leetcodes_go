package main

// https://leetcode.cn/problems/merge-sorted-array/?envType=study-plan-v2&envId=top-interview-150
func merge(nums1 []int, m int, nums2 []int, n int) {

	i, j, p := m-1, n-1, m+n-1
	for {
		if i < 0 {
			if j < 0 || p < 0 {
				break
			} else {
				nums1[p] = nums2[j]
				p--
				j--
			}
		} else if j < 0 {
			if i < 0 || p < 0 {
				break
			} else {
				nums1[p] = nums1[i]
				p--
				i--
			}
		} else if p < 0 {
			break
		} else {
			if nums1[i] < nums2[j] {
				nums1[p] = nums2[j]
				p--
				j--
			} else {
				nums1[p] = nums1[i]
				p--
				i--
			}
		}
	}
}
