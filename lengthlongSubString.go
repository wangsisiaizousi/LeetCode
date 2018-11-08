package main

/*
题目：最长无重复子串
重要思想：
一、滑动窗口
二、边界条件

Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequenceand not a substring.


*/
import "fmt"

func lengthlongSubString(str string) int {
	n := 0
	left := 0
	arr := make([]int, 255)
	for i := 0; i < len(str); i++ {
		if arr[str[i]] == 0 || left > arr[str[i]] {
			n = max(n, i-left+1)
		} else {
			left = arr[str[i]]
		}
		arr[str[i]] = i + 1
	}

	return n
}
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	s := lengthlongSubString("abcabcabab123456")
	fmt.Println(s)
}
