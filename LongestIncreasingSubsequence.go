package main

import "fmt"

//1 4 6 2 8 9 7
func LIS(a[] int)int  {
	l := len(a)
	b:=make([]int,l)
	b[0] = 1
	max:=-1
	result:=0
	for i := 1; i < len(a); i++ {
		//求出0 至 i-1的最大长度max
		for j:=0;j<i;j++{
			if(a[i]>a[j]&&b[j]>max){
				max=b[j]
			}
		}
		b[i] = max+1
		if(result<b[i]){
			result=b[i]
		}

	}
	return result
}
func main() {
	fmt.Println(LIS([]int{1,4,6,2,8,9,7,10,3,11}))
}