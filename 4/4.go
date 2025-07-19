package main

import "fmt"

func main(){
	var n int
	fmt.Scanf("%d\n",&n)
	score := make([]int,n)
	sum:=0
	count:=0
	for i:=0;i<n;i++ {
		fmt.Scanf("%d",&score[i])
		sum+=score[i]
		if score[i]>=90 {
			count++;
		}
	}
	fmt.Printf("sum=%d\n",sum)
	fmt.Printf("count=%d\n",count)
}