/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]

Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.
*/
package main

import (
	"fmt"
)

// Using Array
// Time Complx O(n^2)
func twoSum(nums []int, target int) []int {
	for idx, num := range nums {
		//x := target - num
		for j := idx + 1; j < len(nums); j++ {
			if nums[j] == target-num {
				return []int{idx, j}
			}
		}
	}
	return []int{}
}

// Using HashMap
// Time complx
func twoSum2(nums []int, target int) []int {
	localMap := make(map[int]int) // Key: number, Value: Index of number
	for idx, num := range nums {
		if val, ok := localMap[target-num]; ok {
			return []int{val, idx}
		}
		localMap[num] = idx
	}
	return []int{}
}

func main() {
	fmt.Println("case1: [2,7,11,15], target = 9")
	case1 := []int{2, 7, 11, 15}
	fmt.Printf("Result1: %v\n", twoSum(case1, 9))

	fmt.Println("case2: [3,2,4], target = 6")
	case2 := []int{3, 2, 4}
	fmt.Printf("Result2: %v\n", twoSum(case2, 6))

	fmt.Println("case3: [3,3], target = 6")
	case3 := []int{3, 3}
	fmt.Printf("Result3: %v\n", twoSum(case3, 6))

	fmt.Println("case1_opt2: [2,7,11,15], target = 9")
	case1_opt2 := []int{2, 7, 11, 15}
	fmt.Printf("Result1_opt2: %v\n", twoSum2(case1_opt2, 9))

	fmt.Println("case2_opt2: [3,2,4], target = 6")
	case2_opt2 := []int{3, 2, 4}
	fmt.Printf("Result2_opt2: %v\n", twoSum2(case2_opt2, 6))

	fmt.Println("case3_opt2: [3,3], target = 6")
	case3_opt2 := []int{3, 3}
	fmt.Printf("Result3_opt2: %v\n", twoSum2(case3_opt2, 6))
}
