package main

import (
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
	fmt.Scanf("%d",&n)
	var a []int
	for {
		if n==1 {
			break
		}
		for i:=2;i<=n&&issushu(i);i++ {
			if n%i==0 {
				a=append(a, i)
				n=n/i
				break
			}
		}
	}
	fmt.Println(a)
}