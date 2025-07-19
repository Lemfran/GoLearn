package main

import "fmt"

func binarySearch(num int,arr []int) (int,int){
	left :=0
	right := len(arr)-1
	count :=0

	for left <=right {
		mid:=(left+right)/2
		count++
		if(num == arr[mid]){
			return mid,count
		} else if arr[mid]<num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1,count
}

func main(){
	var n int
	fmt.Scanf("%d\n",&n)
	arr := make([]int,n)
	for i:=0;i<n;i++ {
		fmt.Scanf("%d\n",&arr[i])
	}
	var x int
	fmt.Scanf("%d\n",&x)
	index , count :=binarySearch(x,arr)
	fmt.Println(index)
	fmt.Println(count)
}