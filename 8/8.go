/*
输入n值(1<n<1000)、n个非降序排列的整数以及要查找的数x，使用二分查找算法查找x，输出x所在的下标（0~n-1）及比较次数。 若x不存在，输出-1和比较次数。

输入格式:
输入格式:
第一行是数组的个数n
第二行开始是输入数组的元素
最后一行是需要查找的元素

输出格式:
输出x所在的下标（0~n-1）及比较次数。若x不存在，输出-1和比较次数。

输入样例:
在这里给出一组输入。例如：

4
1 
2
3
4
1
输出样例:
在这里给出相应的输出。例如：

0
2
*/

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