package main

import(
	"fmt"
	"math"
)

func issushu(n int) bool{
	if n<=1 {
		return false
	}
	for i:=2;i<int(math.Sqrt(float64(n)));i++ {
		if n%i==0{
			return false
		}
	}
	return true
	
}

func main(){
	var n int
	fmt.Scanf("%d\n",&n)
	arr :=make([]int,n)

	for i:=0;i<n;i++ {
		fmt.Scanf("%d\n",&arr[i])
	}

	for _,num :=range arr {
		if issushu(num) {
			count:=1
			for i:=3;i<=num;i++{
				if issushu(i){
					count++
				}
			}
			fmt.Printf("%d\n",count)
		} else {
			fmt.Printf("0\n")
		}
	}
}