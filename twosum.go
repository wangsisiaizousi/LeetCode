/*

题目：两数之和
重要思想：
一、来换取时间复杂度

给定一个整数数组，返回两个数字的索引，使它们相加得到一个特定的目标。
假设每个输入都只有一个解决方案，并且可能不会两次使用相同的元素。

Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
Example:
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

*/

package main

import "fmt"

func twoSumSolution(a []int, target int) []int {
	numMap := map[int]int{}
	for i := 0; i < len(a); i++ {
		numMap[a[i]] = i
	}
	for j := 0; j < len(a); j++ {
		num := a[j]
		num = target - num
		value, ok := numMap[num]
		if ok && value != j {
			return []int{j, value}
		}

	}
	return []int{-1, -1}
}

func main() {
	numList := []int{1, 2, 3, 4, 5, 11}
	fmt.Println(&numList[1])
	//ints := append(numList, 21)
	numList[2] = 12
	ints := numList
	fmt.Println(&ints)
	//fmt.Println(twoSumSolution(numList,2))
}
