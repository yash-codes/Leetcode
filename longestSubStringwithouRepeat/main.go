/*
Given a string s, find the length of the longest substring without duplicate characters.



Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct answers.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
//Sliding Window Technique
package main

import "fmt"

func findLengthOfLongestSubstring(s string) int {

	L := 0 //left index of sliding window
	R := 0 //right index of sliding window
	lastSeenMap := make(map[string]int)
	maxLen := 0
	windowLength := 0

	fmt.Printf("Default values: L(%v), R(%v), maxLen(%v), windowLen(%v)\n", L, R, maxLen, windowLength)

	for R = 0; R < len(s); R++ {
		fmt.Printf("\nIteration(%v): L(%v), R(%v), maxLen(%v), windowLen(%v), character(%v)\n", R, L, R, maxLen, windowLength, string(s[R]))
		fmt.Printf("whole String: %v\n", s)
		fmt.Printf("Current window: %v\n", s[L:R+1])
		if lastSeenIdx, exist := lastSeenMap[string(s[R])]; exist && lastSeenIdx >= L {
			fmt.Printf("char(%v) already exist in map(%v), lastSeenIndex for the char(%v) is %v && lastSeenIndex(%v) comes under current Left(%v) & Right(%v) boundaries\n", string(s[R]), lastSeenMap, string(s[R]), lastSeenIdx, lastSeenIdx, L, R)
			L = lastSeenIdx + 1
			fmt.Printf("Shifting Left boundary forward, now Left(%v)\n", L)
		}

		fmt.Printf("Updating Map\n")
		lastSeenMap[string(s[R])] = R
		fmt.Printf("Now Map(%v)\n", lastSeenMap)

		fmt.Printf("Calculating window size..., Before: %v, maxLen: %v\n", windowLength, maxLen)
		windowLength = R - L + 1
		fmt.Printf("Now window size, After: %v\n", windowLength)

		fmt.Printf("checking if window_size(%v) > maxLen(%v)\n", windowLength, maxLen)
		if windowLength > maxLen {
			maxLen = windowLength
			fmt.Printf("if yes updating max size: %v\n", maxLen)
		}
	}
	fmt.Printf("returning maxLength(%v)\n", maxLen)
	return maxLen
}

func findAllLongestSubstring(s string) []string {

	L, R := 0, 0
	maxLen, windowLen := 0, 0
	var subStrings []string
	lastSeenMap := make(map[string]int) // key: character value: last seen index
	longestSubstrings := make(map[string]bool)

	for R = 0; R < len(s); R++ {

		fmt.Printf("\nIteration(%v): L(%v), R(%v), maxLen(%v), windowLen(%v), character(%v)\n", R, L, R, maxLen, windowLen, string(s[R]))
		fmt.Printf("whole String: %v\n", s)
		fmt.Printf("Current window: %v\n", s[L:R+1])

		if lastSeenIdx, exist := lastSeenMap[string(s[R])]; exist && lastSeenIdx >= L {

			fmt.Printf("char(%v) already exist in map(%v), lastSeenIndex for the char(%v) is %v && lastSeenIndex(%v) comes under current Left(%v) & Right(%v) boundaries\n", string(s[R]), lastSeenMap, string(s[R]), lastSeenIdx, lastSeenIdx, L, R)
			L = lastSeenIdx + 1
			//R = R - 1
			fmt.Printf("Shifting Left boundary forward, now Left(%v), Right(%v)\n", L, R)
		}

		fmt.Printf("now Left(%v), Right(%v)\n", L, R)
		fmt.Printf("Updating Map for key(%v):value(%v)\n", string(s[R]), R)
		lastSeenMap[string(s[R])] = R
		fmt.Printf("Now Map(%v)\n", lastSeenMap)

		fmt.Printf("Calculating window size..., Before: %v, maxLen: %v\n", windowLen, maxLen)
		windowLen = R - L + 1

		fmt.Printf("checking if window_size(%v) > maxLen(%v)\n", windowLen, maxLen)
		if windowLen > maxLen {
			maxLen = windowLen
			longestSubstrings = make(map[string]bool)
			fmt.Printf("if yes updating max size: %v\n", maxLen)
		} else if windowLen == maxLen {
			longestSubstrings[s[L-1:R]] = true
			fmt.Printf("longestSubstrings: %v\n", longestSubstrings)
		}
	}

	for key, _ := range longestSubstrings {
		subStrings = append(subStrings, key)
	}
	return subStrings

}

func main() {
	ret1 := findLengthOfLongestSubstring("abcabcbb")
	fmt.Println("Length of longest substring without repeat char: ", ret1)
	fmt.Printf("\n\n")

	ret2 := findAllLongestSubstring("abcabcbb")
	fmt.Println("All longest substring without repeat char: ", ret2)

	ret3 := findAllLongestSubstring("pwwkew")
	fmt.Println("All longest substring without repeat char: ", ret3)

}
