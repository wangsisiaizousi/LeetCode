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
import "fmt"

func lengthlongSubString(str string) int {
	n:=0
	left:=0
	arr:=make([]int,255)
	for i:=0;i<len(str);i++ {
		if(arr[str[i]]==0||left>arr[str[i]]){
			n = max(n,i-left+1)
		}else{
			left = arr[str[i]]
		}
		arr[str[i]] =i+1
	}

	return n
}
func max(a int, b int) int{
	if(a>b){
		return a
	}else {
		return b
	}
}

func main() {
	s:=lengthlongSubString("abcabcabab123456")
	fmt.Println(s)
}