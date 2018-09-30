package main

import "fmt"

func search(nums []int, target int) int {
	pivot := findPivot(nums, 0, len(nums) - 1)

	if pivot == -1 {
		// array is not pivoted
		return binarySearch(nums, 0, len(nums) - 1, target)
	}

	if nums[pivot] == target {
		return pivot
	}

	if nums[0] <= target {
		return binarySearch(nums, 0, pivot - 1, target)
	}
	return binarySearch(nums, pivot + 1, len(nums) - 1, target)
}

func findPivot(nums []int, low int, high int) int {
	if high < low {
		return -1
	}
	if high == low {
		return low
	}

	mid := (low + high)/2
	if mid < high && nums[mid] > nums[mid + 1] {
		return mid
	}
	if mid > low && nums[mid] < nums[mid - 1] {
		return mid - 1
	}
	if nums[low] >= nums[mid] {
		return findPivot(nums, low, mid - 1)
	}
	return findPivot(nums, mid + 1, high)
}

func binarySearch(nums []int, low int, high int, key int) int {
	if high < low {
		return -1
	}

	mid := (low + high) / 2
	if key == nums[mid] {
		return mid
	}

	if key > nums[mid] {
		return binarySearch(nums, mid+1, high, key)
	}
	return binarySearch(nums, low, mid-1, key)
}

func main() {
	result := search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	fmt.Println(result)
}
