package main

import "fmt"

func bruteForceTwoSum(nums []int, target int) []int {
	for idx1 := range nums {
		for idx2 := len(nums) - 1; idx2 > idx1; idx2-- {
			total := nums[idx1] + nums[idx2]
			if total == target {
				return []int{idx1, idx2}
			}
		}
	}
	return []int{}
}

func twoSum(nums []int, target int) []int {
	// Map from value to index.
	lookUp := make(map[int]int)
	for i, val := range nums {
		// If our matching number exists, return indices now.
		j, ok := lookUp[target-val]
		if ok {
			return []int{j, i}
		}
		// Otherwise add to lookup.
		lookUp[val] = i
	}
	return nil
}

func main() {
	inputlist := []int{1, 2, 3, 4, 5}
	target := 9
	fmt.Println(twoSum(inputlist, target))
}
