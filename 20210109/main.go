package main

func main() {

}

func search(n int, nums []int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == n {
			return mid
		} else if n < nums[mid] {
			high = mid - 1
		} else if n > nums[mid] {
			low = mid + 1
		}
	}
	return -1
}
