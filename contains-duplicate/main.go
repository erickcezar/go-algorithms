package main

import "fmt"

func main() {
	nums := []int{1, 2, 8, 9, 2}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	visited := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := visited[nums[i]]; ok {
			return true
		} else {
			visited[nums[i]] = true
		}
	}
	return false
}
