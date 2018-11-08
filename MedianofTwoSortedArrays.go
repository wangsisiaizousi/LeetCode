package main

/*
两个有序数组的中位数

There are two sorted arrays nums1 and nums2 of size m and n respectively.
Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
You may assume nums1 and nums2 cannot be both empty.

Example 1:
nums1 = [1, 3]
nums2 = [2]

The median is 2.0


Example 2:
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5

*/
import (
	"fmt"
	"math"
)

//如果第一个数组的第K/2个数字小的话，那么说明我们要找的数字肯定不在nums1中的前K/2个数字
func findKth(num1 []int, i int, num2 []int, j int, key int) int {
	if i >= len(num1) {
		return j + key - 1
	}
	if j >= len(num2) {
		return i + key - 1
	}
	if key == 1 {
		return minnum(num1[i], num2[j])
	}

	midval1 := 0
	if i+key/2-1 < len(num1) {
		//为什么要减1，因为key/2是num1的中位数，如果比num2的中位数要小的话，下一轮就从中位数开始的
		midval1 = num1[i+key/2-1]
	} else {
		midval1 = math.MaxInt32
	}
	midval2 := 0
	if j+key/2-1 < len(num2) {
		midval2 = num2[j+key/2-1]
	} else {
		midval2 = math.MaxInt32
	}
	if midval2 > midval1 {
		return findKth(num1, i+key/2, num2, j, key-key/2)
	} else {
		return findKth(num1, i, num2, j+key/2, key-key/2)
	}
}
func minnum(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func main() {
	a := []int{1, 3, 5, 7, 9}
	b := []int{2, 4, 6, 8, 10}
	//1 2 3 4 5 6 7 8 9 10
	fmt.Println(findKth(a, 0, b, 0, 6))
}
