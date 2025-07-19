package main

import "fmt"

func main(){
	var n,sum,k int
	fmt.Scanf("%d",&n)
	for i:=0;i<=n;i++ {
		if i==0 {
			k=1
		} else {
			k=k*i
		}
		sum+=k
	}
	fmt.Printf("%d",sum)
}