package main

import (
	"log"
	"testing"
)

type TestCase struct {
	inputList []int
	target    int
}

var testCases = []TestCase{
	TestCase{
		inputList: []int{2, 7, 11, 15},
		target:    9,
	},
	TestCase{
		inputList: []int{2, 5, 5, 11},
		target:    10,
	},
	TestCase{
		inputList: []int{1, 3, 4, 2},
		target:    6,
	},
}

func TestTwoSum(t *testing.T) {
	for _, tc := range testCases {
		indices := twoSum(tc.inputList, tc.target)
		total := tc.inputList[indices[0]] + tc.inputList[indices[1]]
		log.Printf("Found indices: %v, summing to %d", indices, total)
		if total != tc.target {
			t.Errorf("Sum from returned indices (%d) does not match target (%d)", total, tc.target)
		}

	}
}
