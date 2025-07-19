package main

import (
	"fmt"
)

func main(){
	var n int
	fmt.Scanf("%d\n",&n)
	a :=make([]rune,n)
	b :=make([]rune,n)

	for i:=0;i<n;i++{
		fmt.Scanf("%c %c\n",&a[i],&b[i])
	}

	for i:=0;i<n;i++{
		if a[i]==b[i] {
			fmt.Printf("TIE\n")
		} else if (a[i]=='R'&&b[i]=='S')||(a[i]=='S'&&b[i]=='P')||(a[i]=='P'&&b[i]=='R'){
			fmt.Printf("Player1\n")
		} else {
			fmt.Printf("Player2\n")
		}
	}
}