package main

import"fmt"

func main(){
	var n,sum int
	fmt.Scanf("%d",&n)
	for i:=1;i<=n;i++{
		sum+=i*i
	}
	fmt.Printf("%d\n",sum)
}