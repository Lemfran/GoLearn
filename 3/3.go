package main

import "fmt"

func main(){
	var n int
	fmt.Scanf("%d",&n)
	var sum float64
	for i:=1;i<=n;i++{
		if i%2==0 {
			sum-=1/float64(4*i-3)
			continue
		}
		sum+=1/float64(4*i-3)
	}
	fmt.Printf("sum = %.6f",sum)
}